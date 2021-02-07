package senders

import (
	"net/smtp"
	"os"
)

/*
SMTPConfig -> smtp configuration params
*/
type SMTPConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Sender   string
}

/*
NewSMTPConfig -> creates a new SMTPConfig from the given
sender
*/
func NewSMTPConfig(sender string) SMTPConfig {
	return SMTPConfig{
		User:     os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Sender:   sender,
	}
}

/*
EmailSMTPSender -> email smtp sender
*/
type EmailSMTPSender struct {
	config    SMTPConfig
	sendEmail func(string, smtp.Auth, string, []string, []byte) error
}

/*
NewEmailSMTPSender -> creates a new emailSMTPSender from the given
sender
*/
func NewEmailSMTPSender(sender string) EmailSMTPSender {
	return EmailSMTPSender{NewSMTPConfig(sender), smtp.SendMail}
}

/*
Send -> sends the given body to the given receivers
*/
func (e EmailSMTPSender) Send(to []string, body []byte) error {
	addr := e.config.Host + ":" + e.config.Port
	auth := smtp.CRAMMD5Auth(e.config.User, e.config.Password)
	return e.sendEmail(addr, auth, e.config.Sender, to, body)
}
