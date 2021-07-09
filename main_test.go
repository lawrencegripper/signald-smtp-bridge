package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func test_SendEmail_PlainText(t *testing.T) {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "username", "password")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{os.Getenv("SEND_TO") + "@signal.bridge"}
	msg := strings.NewReader("To: recipient@signal.bridge\r\n" +
		"Subject: discount Gophers!\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("localhost:1025", auth, os.Getenv("SEND_FROM")+"@signal.bridge", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func Test_SendEmail_HTML(t *testing.T) {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "username", "password")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{os.Getenv("SEND_TO") + "@signal.bridge"}
	msg := strings.NewReader("To: recipient@signal.bridge\r\n" +
		"Subject: discount Gophers!\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		"<div>hello from div</div>\r\n")
	err := smtp.SendMail("localhost:1025", auth, os.Getenv("SEND_FROM")+"@signal.bridge", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
