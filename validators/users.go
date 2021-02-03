package validators

/*
This module contains structs wich will validate users
notifications input and template data.
*/

/*
EmailUserVerificationValidator -> validator struct
for the /emails/users/verify endpoint
*/
type EmailUserVerificationValidator struct {
	Email            string `json:"email" binding:"required"`
	Name             string `json:"name" binding:"required"`
	VerificationLink string `json:"verification_link" binding:"required"`
}

/*
EmailUserVerificationTemplateData -> user_verification.html
template data struct
*/
type EmailUserVerificationTemplateData struct {
	Name             string `json:"name"`
	VerificationLink string `json:"verification_link"`
}
