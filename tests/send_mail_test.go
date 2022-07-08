package tests

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {

	e := email.NewEmail()
	e.From = fmt.Sprintf("g1aive <%s>", testConfig.Email.Account)
	e.To = []string{testConfig.Email.To}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", testConfig.Email.Account, testConfig.Email.Password, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
