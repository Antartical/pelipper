package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pelipper/models"
	"pelipper/models/senders"
	"pelipper/notices"
	"pelipper/validators"
)

/*
EmailUserVerification -> handler for /email/user/verify.
Sends the user verification email.
*/
func EmailUserVerification(c *gin.Context) {
	var input validators.EmailUserVerificationValidator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := models.Email{
		To:       input.To,
		Subject:  input.Subject,
		Template: "user_verification.html",
		TemplateData: notices.EmailUserVerificationTemplateData{
			Name:             input.Name,
			VerificationLink: input.VerificationLink,
		},
		Sender: senders.NewEmailSMTPSender(input.From),
	}

	if err := email.Deliver(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, nil)
}
