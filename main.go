package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/emersion/go-smtp"
	"github.com/google/uuid"
	"github.com/marcospgmelo/parsemail"
	"gitlab.com/signald/signald-go/signald"
	clientProtocol "gitlab.com/signald/signald-go/signald/client-protocol"
	v1 "gitlab.com/signald/signald-go/signald/client-protocol/v1"
)

var signaldClient *signald.Signald

func init() {
	signaldClient = &signald.Signald{
		SocketPath: "/signald/signald.sock",
	}

	signaldResponses := make(chan clientProtocol.BasicResponse)
	err := signaldClient.Connect()
	if err != nil {
		panic(err)
	}

	go signaldClient.Listen(signaldResponses)
	go func() {
		<-signaldResponses
		panic("Signald closed the channel, usually means an error occurred")
	}()
}

// The Backend implements SMTP server methods.
type Backend struct{}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{Anonymous: true}, nil
}

// A Session is returned after successful login.
type Session struct {
	From        string
	To          string
	MessageData string
	Email       *parsemail.Email
	Anonymous   bool
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	s.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	s.To = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		if os.Getenv("DEBUG") == "TRUE" {
			log.Println("DEBUG Data:", string(b))
		}
		s.MessageData = string(b)
	}

	if err := parseEmail(s); err != nil {
		return err
	}

	return sendSignalMessage(s)
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func (s *Session) AuthPlain(username, password string) error {
	if s.Anonymous {
		if os.Getenv("SMTP_ALLOW_ANNON") == "TRUE" {
			return nil
		} else {
			return smtp.ErrAuthRequired
		}
	}

	if username != os.Getenv("SMTP_USERNAME") && password != os.Getenv("SMTP_PASSWORD") {
		return errors.New("Invalid username or password")
	}

	return nil
}

func main() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func mustGetSignalUserOrGroupFromAddress(address string) string {
	split := strings.Split(address, "@")
	if len(split) < 2 {
		panic("Invalid address must be 'numberOrGroupId@signal.bridge")
	}

	return split[0]
}

func parseEmail(session *Session) error { // this reads an email message
	email, err := parsemail.Parse(strings.NewReader(session.MessageData)) // returns Email struct and error
	if err != nil {
		return err
	}

	log.Println("Subject: " + email.Subject)
	log.Println(email.From)
	log.Println(email.To)

	session.Email = &email

	return nil
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func captureHTMLEmailAsPDF(session *Session) (string, error) {
	// create a test server to serve the page
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = fmt.Fprint(w, session.Email.HTMLBody)
			},
		),
	)
	defer ts.Close()

	// create headless chrome
	resp, err := http.Get("http://localhost:9222/json/version")
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), result["webSocketDebuggerUrl"].(string))
	defer cancelActxt()
	ctx, cancel := chromedp.NewContext(actxt)
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(ts.URL, &buf)); err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("/signald/%semail.pdf", uuid.New().String())
	log.Printf("PDF using file: %q", filePath)
	if err := ioutil.WriteFile(filePath, buf, 0777); err != nil {
		return "", err
	}
	return filePath, nil
}

var phoneNumberRegex, _ = regexp.Compile("\\+?44[0-9]{10}")

func sendSignalMessage(session *Session) error {
	var pdfFile string
	var err error
	if session.Email.ContentType != "text/plain" {
		log.Println("Converting HTML mail to pdf file")
		pdfFile, err = captureHTMLEmailAsPDF(session)
		if err != nil {
			log.Println("PDF conversion failed")
		}
	}

	signalMsg := session.From + "\n\n" + session.Email.Subject + "\n\n" + session.Email.TextBody

	var fromUsername string
	if strings.Contains(session.From, "@signal.bridge") {
		if !strings.HasPrefix(session.From, "+") {
			fromUsername = os.Getenv("SEND_FROM")
		} else {
			fromUsername = mustGetSignalUserOrGroupFromAddress(session.From)
		}
	} else {
		fromUsername = os.Getenv("SEND_FROM")
	}

	log.Printf("Converting email session to signal msg")
	log.Printf("Sending from account: %q", fromUsername)

	req := v1.SendRequest{
		Username:    fromUsername,
		MessageBody: signalMsg,
	}

	// check file exists
	_, err = os.Stat(pdfFile)
	if err == nil {
		req.Attachments = []*v1.JsonAttachment{
			{Filename: pdfFile},
		}
	}

	if strings.Contains(session.To, "@signal.bridge") {
		recipient := mustGetSignalUserOrGroupFromAddress(session.To)

		if phoneNumberRegex.MatchString(recipient) {
			if !strings.HasPrefix(recipient, "+") {
				recipient = "+" + recipient
			}
			req.RecipientAddress = &v1.JsonAddress{Number: recipient}
			log.Printf("Sending to user: %q", recipient)
		} else {
			req.RecipientGroupID = recipient
			log.Printf("Sending to group: %q", recipient)
		}
	} else {
		req.RecipientAddress = &v1.JsonAddress{Number: os.Getenv("SEND_TO")}
	}

	if os.Getenv("DEBUG") == "TRUE" {
		log.Printf("DEBUG signal send request: %+v", req)
	}

	resp, err := req.Submit(signaldClient)
	if err != nil {
		log.Printf("crashing -> error sending request to signald: %+v\n", err)
		os.Exit(1)
	}
	for _, msgSent := range resp.Results {
		log.Printf("Sent to: %s in %v ms\n", msgSent.Address.Number, msgSent.Success.Duration)
	}

	_, err = os.Stat(pdfFile)
	if err == nil {
		os.Remove(pdfFile)
	}

	return nil
}
