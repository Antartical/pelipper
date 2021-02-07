package senders

import (
	"net/smtp"
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
		sender := "test@test.com"
		smtpConfig := NewSMTPConfig(sender)

		assert.Equal(smtpConfig.Sender, sender)
	})
}

func TestEmailSMTPSender(t *testing.T) {
	assert := require.New(t)

	t.Run("Test EmailSMTPSender constructor", func(t *testing.T) {
		sender := "test@test.com"
		emailSMTPSender := NewEmailSMTPSender(sender)

		assert.Equal(emailSMTPSender.config.Sender, sender)
	})

	t.Run("Test EmailSMTPSender Send", func(t *testing.T) {
		to := []string{"test@test.com"}
		from := "from-test@test.com"
		body := []byte("Hello")
		mockSendEmail, recorder := mockSend(nil)
		sender := EmailSMTPSender{
			config:    NewSMTPConfig(from),
			sendEmail: mockSendEmail,
		}

		expectedAddr := sender.config.Host + ":" + sender.config.Port
		expectedAuth := smtp.CRAMMD5Auth(sender.config.User, sender.config.Password)

		sender.Send(to, body)

		assert.Equal(recorder.addr, expectedAddr)
		assert.Equal(recorder.auth, expectedAuth)
		assert.Equal(recorder.from, from)
		assert.Equal(recorder.to, to)
		assert.Equal(recorder.body, body)

	})
}
