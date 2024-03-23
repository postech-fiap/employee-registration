package repository

import (
	"net/smtp"

	"github.com/postech-fiap/employee-registration/cmd/config"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

type emailRepository struct {
	config *config.Config
}

func NewEmailRepository(config *config.Config) port.EmailRepository {
	return emailRepository{config: config}
}

func (r emailRepository) SendEmail(to, subject, body string) error {
	contet := `From: Employee Resgistration <` + r.config.SMTP.From + `>
To: Colaborador <` + to + `>
Subject: ` + subject + `
Content-Type: multipart/alternative; boundary="boundary-string"

--boundary-string
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline

` + subject + `

--boundary-string
Content-Type: text/html; charset="utf-8"
Content-Transfer-Encoding: quoted-printable
Content-Disposition: inline
<!doctype html>
<html>
  <body>
` + body + `
</body>
</html>
--boundary-string--
`
	return smtp.SendMail(
		r.config.SMTP.Host+":"+r.config.SMTP.Port,
		smtp.PlainAuth(
			"",
			r.config.SMTP.Username,
			r.config.SMTP.Password,
			r.config.SMTP.Host,
		),
		r.config.SMTP.From,
		[]string{to},
		[]byte(contet),
	)
}
