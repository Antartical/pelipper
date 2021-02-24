package notices

/*
EmailUserVerificationTemplateData -> user_verification.html
template data struct
*/
type EmailUserVerificationTemplateData struct {
	Name             string
	VerificationLink string
}

/*
EmailUserChangePasswordTemplateData -> user_change_password_html.html
template data struct
*/
type EmailUserChangePasswordTemplateData struct {
	Name               string
	ChangePasswordLink string
}
