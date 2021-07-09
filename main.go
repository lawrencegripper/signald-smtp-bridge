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
	"strings"
	"time"

	"github.com/DusanKasan/parsemail"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/emersion/go-smtp"
	"gitlab.com/signald/signald-go/signald"
	v0 "gitlab.com/signald/signald-go/signald/client-protocol/v0"
	v1 "gitlab.com/signald/signald-go/signald/client-protocol/v1"
)

var signaldClient *signald.Signald

func init() {
	signaldClient = &signald.Signald{
		SocketPath: "/signald/signald.sock",
	}
	err := signaldClient.Connect()
	if err != nil {
		panic(err)
	}
	go signaldClient.Listen(nil)
}

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != os.Getenv("SMTP_USERNAME") || password != os.Getenv("SMTP_PASSWORD") {
		return nil, errors.New("Invalid username or password")
	}
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	if os.Getenv("SMTP_ALLOW_ANNON") == "TRUE" {
		return &Session{}, nil
	}
	return nil, smtp.ErrAuthRequired
}

// A Session is returned after successful login.
type Session struct {
	From        string
	To          string
	MessageData string
	Email       *parsemail.Email
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
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
		log.Println("Data:", string(b))
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

func mustGetRecipientFromAddress(address string) string {
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

	log.Println(email.Subject)
	log.Println(email.From)
	log.Println(email.To)
	log.Println(email.HTMLBody)

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

func captureHtmlEmailAsPDF(session *Session) error {
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
		return err
	}

	if err := ioutil.WriteFile("/signald/email.pdf", buf, 0777); err != nil {
		return err
	}
	return nil
}

func sendSignalMessage(session *Session) error {
	log.Printf("Converting email session to signal msg: %+v", session)

	if session.Email.ContentType != "text/plain" {
		captureHtmlEmailAsPDF(session)
	}

	signalMsg := session.Email.Subject + "\n\n" + session.Email.TextBody

	req := v1.SendRequest{
		Username:    mustGetRecipientFromAddress(session.From),
		MessageBody: signalMsg,
	}

	// check file exists
	_, err := os.Stat("/signald/email.pdf")
	if err == nil {
		req.Attachments = []*v0.JsonAttachment{
			{Filename: "/signald/email.pdf"},
		}
	}

	recipient := mustGetRecipientFromAddress(session.To)
	if strings.HasPrefix(recipient, "+") {
		req.RecipientAddress = &v1.JsonAddress{Number: recipient}
	} else {
		req.RecipientGroupID = recipient
	}

	log.Printf("Sending message with signal: %+v", req)
	resp, err := req.Submit(signaldClient)
	if err != nil {
		log.Printf("error sending request to signald: %+v\n", err)
	}
	for _, msgSent := range resp.Results {
		log.Printf("Sent to: %s in %v ms\n", msgSent.Address.Number, msgSent.Success.Duration)
	}

	return nil
}
