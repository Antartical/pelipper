package services

import (
	"pelipper/models"
	"pelipper/models/senders"
)

/*
IEmailService -> interface for possibles EmailServices
*/
type IEmailService interface {
	SendEmail(from string, to string, subject string, template string, templateData interface{}) error
}

/*
EmailService -> implements IEmailService and it is the one who will manage the
emails deliveries-
*/
type EmailService struct {
	sender models.IEmailSender
}

/*
NewSMTPEmailService -> creates an emailService instance which will use
SMTP server to send messages
*/
func NewSMTPEmailService() EmailService {
	return EmailService{
		sender: senders.NewEmailSMTPSender(),
	}
}

/*
SendEmail -> sends an email with the given params
*/
func (e EmailService) SendEmail(from string, to string, subject string, template string, templateData interface{}) error {
	email := models.Email{
		From:         from,
		To:           to,
		Subject:      subject,
		Template:     template,
		TemplateData: templateData,
		Sender:       e.sender}
	return email.Deliver()
}
