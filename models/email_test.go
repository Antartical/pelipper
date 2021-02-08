package models

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pelipper/notices"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

type emailSenderMock struct {
	send func(to []string, body []byte) error
}

func (e emailSenderMock) Send(to []string, body []byte) error {
	return e.send(to, body)
}

type emailSendRecorder struct {
	to   []string
	body []byte
}

func mockSend(err error) (func(to []string, body []byte) error, *emailSendRecorder) {
	r := new(emailSendRecorder)
	return func(to []string, body []byte) error {
		*r = emailSendRecorder{to, body}
		return err
	}, r
}

func TestEmailDeliver(t *testing.T) {

	assert := require.New(t)
	to := "test@test.com"
	subject := "test"
	templatesDir := os.Getenv("TEMPLATES_DIR")
	templateSuccess := "user_verification.html"
	templateError := "unknown.html"
	templateSuccessData := notices.EmailUserVerificationTemplateData{
		Name:             "test",
		VerificationLink: "http://test.com",
	}

	t.Run("Test email deliver successfully", func(t *testing.T) {
		mockSend, sendRecorder := mockSend(nil)
		email := Email{
			To:           to,
			Subject:      subject,
			Template:     templateSuccess,
			TemplateData: templateSuccessData,
			Sender:       emailSenderMock{mockSend},
		}

		var expectedBody bytes.Buffer
		expectedBody.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))
		data, _ := template.ParseFiles(
			filepath.Join(templatesDir, templateSuccess),
		)
		data.Execute(&expectedBody, templateSuccessData)

		email.Deliver()
		assert.Equal(sendRecorder.to, []string{to})
		assert.Equal(sendRecorder.body, expectedBody.Bytes())
	})

	t.Run("Test email deliver with template error", func(t *testing.T) {

		mockSend, _ := mockSend(nil)
		email := Email{
			To:           to,
			Subject:      subject,
			Template:     templateError,
			TemplateData: templateSuccessData,
			Sender:       emailSenderMock{mockSend},
		}

		assert.NotNil(email.Deliver())
	})

	t.Run("Test email deliver SMTP error", func(t *testing.T) {

		mockSend, _ := mockSend(errors.New("Test error"))
		email := Email{
			To:           to,
			Subject:      subject,
			Template:     templateSuccess,
			TemplateData: templateSuccessData,
			Sender:       emailSenderMock{mockSend},
		}
		assert.NotNil(email.Deliver())
	})
}
