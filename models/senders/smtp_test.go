package senders

import (
	"net/smtp"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type sendEmailMockRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	body []byte
}

func mockSend(err error) (func(string, smtp.Auth, string, []string, []byte) error, *sendEmailMockRecorder) {
	r := new(sendEmailMockRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, body []byte) error {
		*r = sendEmailMockRecorder{addr, a, from, to, body}
		return err
	}, r
}

func TestSMTPConfig(t *testing.T) {
	assert := require.New(t)
	t.Run("Test SMTPConfig constructor", func(t *testing.T) {
		smtpConfig := NewSMTPConfig()
		assert.Equal(smtpConfig.Host, os.Getenv("SMTP_HOST"))
	})
}

func TestEmailSMTPSender(t *testing.T) {
	assert := require.New(t)

	t.Run("Test EmailSMTPSender constructor", func(t *testing.T) {
		emailSMTPSender := NewEmailSMTPSender()
		assert.Equal(emailSMTPSender.config.Host, os.Getenv("SMTP_HOST"))
	})

	t.Run("Test EmailSMTPSender Send", func(t *testing.T) {
		assert.Fail("")
		to := []string{"test@test.com"}
		from := "from-test@test.com"
		body := []byte("Hello")
		mockSendEmail, recorder := mockSend(nil)
		sender := EmailSMTPSender{
			config:    NewSMTPConfig(),
			sendEmail: mockSendEmail,
		}

		expectedAddr := sender.config.Host + ":" + sender.config.Port
		expectedAuth := smtp.CRAMMD5Auth(sender.config.User, sender.config.Password)

		sender.Send(from, to, body)

		assert.Equal(recorder.addr, expectedAddr)
		assert.Equal(recorder.auth, expectedAuth)
		assert.Equal(recorder.from, from)
		assert.Equal(recorder.to, to)
		assert.Equal(recorder.body, body)

	})
}
