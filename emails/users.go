package mails

/*
EmailUserVerificationTemplateData -> user_verification.html
template data struct
*/
type EmailUserVerificationTemplateData struct {
	Name             string
	VerificationLink string
}
