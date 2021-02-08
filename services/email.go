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
SMTPEmailService -> implements IEmailService and it is the one who will manage the
emails deliveries-
*/
type SMTPEmailService struct{}

/*
SendEmail -> sends an email with the given params
*/
func (e SMTPEmailService) SendEmail(from string, to string, subject string, template string, templateData interface{}) error {
	email := models.Email{
		To:           to,
		Subject:      subject,
		Template:     template,
		TemplateData: templateData,
		Sender:       senders.NewEmailSMTPSender(from)}
	return email.Deliver()
}
