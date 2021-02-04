package validators

/*
EmailUserVerificationValidator -> validator struct
for the /emails/users/verify endpoint
*/
type EmailUserVerificationValidator struct {
	From             string `json:"from" binding:"required"`
	To               string `to:"email" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Subject          string `json:"subject" binding:"required"`
	VerificationLink string `json:"verification_link" binding:"required"`
}
