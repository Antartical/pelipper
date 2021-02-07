package models

import (
	"errors"
	"net/smtp"
	"os"
	"pelipper/notices"
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

func TestEmailDeliver(t *testing.T) {

	assert := require.New(t)
	to := "test@test.com"
	from := "http://test.com"
	addr := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	auth := smtp.CRAMMD5Auth(
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)

	t.Run("Test email deliver successfully", func(t *testing.T) {
		mockFunction, mockRegister := mockSend(nil)
		email := Email{
			To:       to,
			Subject:  "test",
			Template: "user_verification.html",
			TemplateData: notices.EmailUserVerificationTemplateData{
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
	})

	t.Run("Test email deliver with template error", func(t *testing.T) {
		assert := require.New(t)

		email := Email{
			To:       to,
			Subject:  "test",
			Template: "unknown.html",
			TemplateData: notices.EmailUserVerificationTemplateData{
				Name:             "test",
				VerificationLink: "http://test.com",
			},
			Sender: NewEmailSMTPSender(from),
		}

		assert.NotNil(email.Deliver())
	})

	t.Run("Test email deliver SMTP error", func(t *testing.T) {
		assert := require.New(t)

		mockFunction, _ := mockSend(errors.New("Test error"))
		email := Email{
			To:       to,
			Subject:  "test",
			Template: "user_verification.html",
			TemplateData: notices.EmailUserVerificationTemplateData{
				Name:             "test",
				VerificationLink: "http://test.com",
			},
			Sender: EmailSMTPSender{
				config: NewSMTPConfig(from),
				send:   mockFunction,
			},
		}

		assert.NotNil(email.Deliver())
	})
}
