package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pelipper/notices"
	"pelipper/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type emailServiceMock struct {
	recorder *sendEmailRecorderMock
	err      error
}

func (e emailServiceMock) SendEmail(from string, to string, subject string, template string, templateData interface{}) error {
	*e.recorder = sendEmailRecorderMock{from, to, subject, template, templateData}
	return e.err
}

func newEmailServiceMock(err error) emailServiceMock {
	return emailServiceMock{new(sendEmailRecorderMock), err}
}

type sendEmailRecorderMock struct {
	From         string
	To           string
	Subject      string
	Template     string
	TemplateData interface{}
}

func setupUsersRouter(emailService services.IEmailService) *gin.Engine {
	router := gin.Default()
	RegisterUserRoutes(router, emailService)
	return router
}

func TestEmailUserVerify(t *testing.T) {
	assert := require.New(t)
	var response gin.H

	t.Run("Test send email user successfully", func(t *testing.T) {
		emailService := newEmailServiceMock(nil)
		router := setupUsersRouter(emailService)

		from := "test@test.com"
		to := "test-receiver@test.com"
		subject := "Test email"
		name := "Test misco"
		verificationLink := "http://test.test"
		templateData := notices.EmailUserVerificationTemplateData{
			Name:             name,
			VerificationLink: verificationLink,
		}

		payload, _ := json.Marshal(map[string]string{
			"from":              from,
			"to":                to,
			"subject":           subject,
			"Name":              name,
			"verification_link": verificationLink,
		})
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(
			"POST", "/emails/users/verify", bytes.NewBuffer(payload),
		)
		router.ServeHTTP(recorder, request)

		err := json.Unmarshal(recorder.Body.Bytes(), &response)
		if err != nil {
			assert.Fail("Payload does not match with the expected one")
		}

		assert.Equal(recorder.Result().StatusCode, http.StatusCreated)
		assert.Nil(response)
		assert.Equal(emailService.recorder.From, from)
		assert.Equal(emailService.recorder.To, to)
		assert.Equal(emailService.recorder.Subject, subject)
		assert.Equal(emailService.recorder.Template, "user_verification.html")
		assert.Equal(emailService.recorder.TemplateData, templateData)
	})

	t.Run("Test send email bad request", func(t *testing.T) {

	})
}
