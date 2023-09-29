package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
	FromEmailPassword = "dwxljdwksncaunsy"
	Name              = "Nillavee cakes and Patries"
	FromEmailAddress  = "myrachanto1@gmail.com"
)

type SenderInterface interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bc []string,
		attachments []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) SenderInterface {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}
func (g *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bc []string, attachments []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bc
	for _, f := range attachments {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attache this file %s with %s", f, err)
		}
	}
	smtpAuth := smtp.PlainAuth("", g.fromEmailAddress, g.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}

func main() {

}
