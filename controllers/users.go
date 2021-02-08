package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pelipper/notices"
	"pelipper/services"
	"pelipper/validators"
)

/*
RegisterUserRoutes -> register user endpoints to the given router
*/
func RegisterUserRoutes(router *gin.Engine, emailService services.IEmailService) {
	users := router.Group("/emails/users")
	users.Use(func(c *gin.Context) {
		c.Set("emailService", emailService)
	})
	users.POST("/verify", EmailUserVerification)
}

/*
EmailUserVerification -> handler for /email/user/verify.
Sends the user verification email.
*/
func EmailUserVerification(c *gin.Context) {
	emailService := c.MustGet("emailService").(services.IEmailService)

	var input validators.EmailUserVerificationValidator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template := "user_verification.html"
	templateData := notices.EmailUserVerificationTemplateData{
		Name:             input.Name,
		VerificationLink: input.VerificationLink,
	}

	err := emailService.SendEmail(
		input.From, input.To, input.Subject, template, templateData,
	)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, nil)
}
