package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
}
