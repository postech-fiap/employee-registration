package port

type EmailRepository interface {
	SendEmail(to, subject, body string) error
}
