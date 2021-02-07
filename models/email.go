package models

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

/*
EmailSender -> interface for email delivery
*/
type EmailSender interface {
	Send(to []string, body []byte) error
}

/*
Email -> email struct
*/
type Email struct {
	To           string
	Subject      string
	Template     string
	TemplateData interface{}
	Sender       EmailSender
}

/*
Deliver -> send email
*/
func (e Email) Deliver() error {
	receivers := []string{e.To}
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", e.Subject, mimeHeaders)))
	t, err := template.ParseFiles(
		filepath.Join(os.Getenv("TEMPLATES_DIR"), e.Template),
	)
	if err != nil {
		return err
	}
	t.Execute(&body, e.TemplateData)
	return e.Sender.Send(receivers, body.Bytes())
}
