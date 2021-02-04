package mails

/*
EmailUserVerificationTemplateData -> user_verification.html
template data struct
*/
type EmailUserVerificationTemplateData struct {
	Name             string `json:"name"`
	VerificationLink string `json:"verification_link"`
}
