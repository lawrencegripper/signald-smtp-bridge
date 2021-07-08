package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"gitlab.com/signald/signald-go/signald"
	v1 "gitlab.com/signald/signald-go/signald/client-protocol/v1"
)

var signaldClient *signald.Signald
var numberExtraction *regexp.Regexp

func init() {
	signaldClient = &signald.Signald{
		SocketPath: "/var/run/signald/signald.sock",
	}
	err := signaldClient.Connect()
	if err != nil {
		panic(err)
	}
	go signaldClient.Listen(nil)

	numberExtraction = regexp.MustCompile(`\+[0-9]*`)
}

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != "username" || password != "password" {
		return nil, errors.New("Invalid username or password")
	}
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

// A Session is returned after successful login.
type Session struct {
	From        string
	To          string
	MessageBody string
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
		s.MessageBody = string(b)
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

	s.Addr = ":25"
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

func mustGetNumberFromAddress(address string) string {
	return numberExtraction.FindString(address)
}

func sendSignalMessage(session *Session) error {
	log.Printf("Sending message with signal: %v", session)
	req := v1.SendRequest{
		Username:    mustGetNumberFromAddress(session.From),
		MessageBody: session.MessageBody,
	}

	if strings.HasPrefix(mustGetNumberFromAddress(session.To), "+") {
		req.RecipientAddress = &v1.JsonAddress{Number: session.To}
	} else {
		// req.RecipientGroupID = args[0]
		panic("not implemented")
	}

	resp, err := req.Submit(signaldClient)
	if err != nil {
		log.Printf("error sending request to signald: %v\n", err)
	}
	log.Printf("Response: %v \n", resp)

	return nil
}
