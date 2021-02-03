package helpers

/*
This module contains methods that will help us with the
email delivery.
*/

import (
	"bytes"
	"fmt"
	templates "html/template"
	"net/smtp"
	"os"
	"path/filepath"
)

const mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

/*
SendEmail -> send an email with the given params
*/
func SendEmail(from string, to string, subject string, template string, data interface{}) error {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	auth := smtp.CRAMMD5Auth(
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_PASSWORD"),
	)

	receivers := []string{to}
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf(
		"Subject: %s\n%s\n\n",
		subject,
		mimeHeaders,
	)))

	t, err := templates.ParseFiles(filepath.Join(os.Getenv("TEMPLATES_DIR"), template))

	if err != nil {
		return err
	}
	t.Execute(&body, data)

	return smtp.SendMail(addr, auth, from, receivers, body.Bytes())
}
