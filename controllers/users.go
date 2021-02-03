package controllers

/*
This module contains handlers for users notifications
*/

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	helpers "pelipper/helpers"
	validators "pelipper/validators"
)

/*
EmailUserVerification -> handler for /email/user/verify

This endpoint will send the verification link to the given
email. The payload must match the following schema:

{
	"email": "",
	"name": "",
	"verification_link": ""
}
*/
func EmailUserVerification(c *gin.Context) {
	var input validators.EmailUserVerificationValidator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	templateData := validators.EmailUserVerificationTemplateData{
		Name:             input.Name,
		VerificationLink: input.VerificationLink,
	}
	err := helpers.SendEmail(
		os.Getenv("HODOR_SENDER"),
		input.Email,
		"Verify your account",
		"hodor/user_verification.html",
		templateData,
	)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, nil)
}
