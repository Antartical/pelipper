package models

import (
	"net/smtp"
	"os"
	"pelipper/mails"
	"testing"

	"github.com/stretchr/testify/require"
)

type EmailSMTPSenderMock struct {
	config SMTPConfig
	send   func(string, smtp.Auth, string, []string, []byte) error
}

type emailMockRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}

func mockSend(err error) (func(string, smtp.Auth, string, []string, []byte) error, *emailMockRecorder) {
	r := new(emailMockRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailMockRecorder{addr, a, from, to, msg}
		return err
	}, r
}

func TestEmailDeliverSuccessfully(t *testing.T) {
	assert := require.New(t)

	to := "test@test.com"
	from := "http://test.com"
	addr := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	auth := smtp.CRAMMD5Auth(
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)

	mockFunction, mockRegister := mockSend(nil)
	email := Email{
		To:       to,
		Subject:  "test",
		Template: "user_verification.html",
		TemplateData: mails.EmailUserVerificationTemplateData{
			Name:             "test",
			VerificationLink: "http://test.com",
		},
		Sender: EmailSMTPSender{
			config: NewSMTPConfig(from),
			send:   mockFunction,
		},
	}

	email.Deliver()
	assert.Equal([]string{to}, mockRegister.to)
	assert.Equal(from, mockRegister.from)
	assert.Equal(addr, mockRegister.addr)
	assert.Equal(auth, mockRegister.auth)
}

func TestEmailDeliverTemplateError(t *testing.T) {
	assert := require.New(t)

	email := Email{
		To:       "test",
		Subject:  "test",
		Template: "user_verification.html",
		TemplateData: mails.EmailUserVerificationTemplateData{
			Name:             "test",
			VerificationLink: "http://test.com",
		},
		Sender: NewEmailSMTPSender("from@from.com"),
	}

	assert.Nil(email.Deliver())

}

func TestEmailDeliverSMTPError(t *testing.T) {
	assert := require.New(t)

	mockFunction, _ := mockSend(nil)
	email := Email{
		To:       "test@test.com",
		Subject:  "test",
		Template: "user_verification.html",
		TemplateData: mails.EmailUserVerificationTemplateData{
			Name:             "test",
			VerificationLink: "http://test.com",
		},
		Sender: EmailSMTPSender{
			config: NewSMTPConfig("http://test.com"),
			send:   mockFunction,
		},
	}

	assert.Nil(email.Deliver())
}
