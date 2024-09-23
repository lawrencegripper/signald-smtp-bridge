package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/mail"
	"os"
	"regexp"
	"strings"
	"time"

	"io/fs"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/emersion/go-smtp"
	"github.com/google/uuid"
	"github.com/matrix-org/mautrix-go"
	"github.com/matrix-org/mautrix-go/mautrix"
	"github.com/matrix-org/mautrix-go/mautrix/event"
	"github.com/matrix-org/mautrix-go/mautrix/id"
)

var matrixClient *mautrix.Client

func init() {
	serverURL := os.Getenv("SYNAPSE_SERVER_URL")
	token := os.Getenv("SYNAPSE_TOKEN")

	var err error
	matrixClient, err = mautrix.NewClient(serverURL, "", token)
	if err != nil {
		panic(err)
	}
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
	Body        string
	Subject     string
	ContentType string
	Anonymous   bool
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	// s.From = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	// s.To = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if bodyBytes, err := io.ReadAll(r); err != nil {
		log.Println(err)
		return err
	} else {
		if os.Getenv("DEBUG") == "TRUE" {
			log.Println("DEBUG Data:", string(bodyBytes))
		}
		s.MessageData = string(bodyBytes)
	}

	if err := parseEmail(s); err != nil {
		log.Println(err)
		return err
	}

	return sendMatrixMessage(s)
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

func mustGetMatrixUserOrGroupFromAddress(address string) string {
	split := strings.Split(address, "@")
	if len(split) < 2 {
		panic("Invalid address must be 'userOrGroupId@matrix.bridge")
	}

	return split[0]
}

func parseEmail(session *Session) error { // this reads an email message
	email, err := mail.ReadMessage(strings.NewReader(session.MessageData)) // returns Email struct and error
	if err != nil {
		return err
	}

	log.Println("Subject: " + email.Header.Get("Subject"))
	session.From = email.Header.Get("From")
	session.To = email.Header.Get("To")
	session.Subject = email.Header.Get("Subject")
	session.ContentType = email.Header.Get("Content-type")
	bodyBytes, err := io.ReadAll(email.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
	}
	session.Body = string(bodyBytes)

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
				_, _ = fmt.Fprint(w, session.Body)
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
	if err := os.WriteFile(filePath, buf, fs.FileMode(0777)); err != nil {
		return "", err
	}
	return filePath, nil
}

var phoneNumberRegex, _ = regexp.Compile("\\+?44[0-9]{10}")

func sendMatrixMessage(session *Session) error {
	var pdfFile string
	var err error
	if session.ContentType != "text/plain" {
		log.Println("Converting HTML mail to pdf file")
		pdfFile, err = captureHTMLEmailAsPDF(session)
		if err != nil {
			log.Println("PDF conversion failed")
		}
	}

	matrixMsg := session.From + "\n\n" + session.Subject + "\n\n" + session.Body

	var fromUsername string
	if strings.Contains(session.From, "@matrix.bridge") {
		if !strings.HasPrefix(session.From, "@") {
			fromUsername = os.Getenv("SEND_FROM")
		} else {
			fromUsername = mustGetMatrixUserOrGroupFromAddress(session.From)
		}
	} else {
		fromUsername = os.Getenv("SEND_FROM")
	}

	log.Printf("Converting email session to matrix msg")
	log.Printf("Sending from account: %q", fromUsername)

	content := event.MessageEventContent{
		MsgType: event.MsgText,
		Body:    matrixMsg,
	}

	if strings.Contains(session.To, "@matrix.bridge") {
		recipient := mustGetMatrixUserOrGroupFromAddress(session.To)

		if phoneNumberRegex.MatchString(recipient) {
			if !strings.HasPrefix(recipient, "@") {
				recipient = "@" + recipient
			}
			_, err = matrixClient.SendMessageEvent(id.RoomID(recipient), event.EventMessage, content)
			if err != nil {
				log.Printf("Error sending message to user: %q", recipient)
			} else {
				log.Printf("Sent to user: %q", recipient)
			}
		} else {
			_, err = matrixClient.SendMessageEvent(id.RoomID(recipient), event.EventMessage, content)
			if err != nil {
				log.Printf("Error sending message to group: %q", recipient)
			} else {
				log.Printf("Sent to group: %q", recipient)
			}
		}
	} else {
		_, err = matrixClient.SendMessageEvent(id.RoomID(os.Getenv("SEND_TO")), event.EventMessage, content)
		if err != nil {
			log.Printf("Error sending message to default recipient: %q", os.Getenv("SEND_TO"))
		} else {
			log.Printf("Sent to default recipient: %q", os.Getenv("SEND_TO"))
		}
	}

	if pdfFile != "" {
		os.Remove(pdfFile)
	}

	return nil
}
