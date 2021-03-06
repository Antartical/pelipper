package services

import (
	"pelipper/models/senders"
	"pelipper/notices"
	"testing"

	"github.com/stretchr/testify/require"
)

type emailSenderMock struct {
	send func(from string, to []string, body []byte) error
}

func (e emailSenderMock) Send(from string, to []string, body []byte) error {
	return e.send(from, to, body)
}

type emailSendRecorder struct {
	from string
	to   []string
	body []byte
}

func mockSend(err error) (func(from string, to []string, body []byte) error, *emailSendRecorder) {
	r := new(emailSendRecorder)
	return func(from string, to []string, body []byte) error {
		*r = emailSendRecorder{from, to, body}
		return err
	}, r
}

func TestSMTPEmailService(t *testing.T) {
	assert := require.New(t)
	from := "test@test.com"
	to := "test@test.com"
	subject := "test"
	templateSuccessData := notices.EmailUserVerificationTemplateData{
		Name:             "test",
		VerificationLink: "http://test.com",
	}

	t.Run("Test SMTP email service constructor", func(t *testing.T) {
		emailService := NewSMTPEmailService()
		assert.IsType(emailService.sender, senders.NewEmailSMTPSender())
	})

	t.Run("Test SendEmail", func(t *testing.T) {
		mockSendF, sendRecorder := mockSend(nil)
		emailService := EmailService{sender: emailSenderMock{mockSendF}}
		emailService.SendEmail(from, to, subject, "user_verification.html", templateSuccessData)

		assert.Equal(sendRecorder.from, from)
		assert.Equal(sendRecorder.to, []string{to})
	})
}
