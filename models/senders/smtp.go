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
}

/*
NewSMTPConfig -> creates a new SMTPConfig from the given
sender
*/
func NewSMTPConfig() SMTPConfig {
	return SMTPConfig{
		User:     os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
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
func NewEmailSMTPSender() EmailSMTPSender {
	return EmailSMTPSender{NewSMTPConfig(), smtp.SendMail}
}

/*
Send -> sends the given body to the given receivers
*/
func (e EmailSMTPSender) Send(from string, to []string, body []byte) error {
	addr := e.config.Host + ":" + e.config.Port
	auth := smtp.CRAMMD5Auth(e.config.User, e.config.Password)
	return e.sendEmail(addr, auth, from, to, body)
}
