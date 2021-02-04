package models

/*
EmailSender -> interface for email delivery
*/
type EmailSender interface {
	Send(to []string, body []byte) error
}
