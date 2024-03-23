package usecase

import (
	"fmt"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type reportUseCase struct {
	mirrorUseCase   port.MirrorUseCase
	emailRepository port.EmailRepository
}

func NewReportUseCase(mirrorUseCase port.MirrorUseCase, emailRepository port.EmailRepository) port.ReportUseCase {
	return reportUseCase{
		mirrorUseCase:   mirrorUseCase,
		emailRepository: emailRepository,
	}
}

func (e reportUseCase) GenarateMirrorReport(userId int) error {
	mirror, err := e.mirrorUseCase.GetMirror(userId)
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("Espelho de ponto - %s/%v", mirror.Month, mirror.Year)

	if err = e.emailRepository.SendEmail(mirror.Email, subject, generateBody(*mirror)); err != nil {
		return err
	}

	return nil
}

func generateBody(mirror domain.Mirror) string {
	body := `Olá %s, seu espelho de ponto chegou!<br />Referente ao mês de <b>%s/%v</b><hr /><b>Data &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Batidas</b><hr />`
	var before time.Time
	for _, register := range mirror.Registers {
		if before.IsZero() || before.Day() != register.Day() {
			body += "<br /><b>" + register.Format("02/01/2006") + "</b>"
		}
		body += `&nbsp;&nbsp;&nbsp;` + register.Format("15:04")
		before = register
	}

	body += `<hr />Horas trabalhadas: %s`

	return fmt.Sprintf(body,
		mirror.Name,
		mirror.Month,
		mirror.Year,
		mirror.Hours,
	)
}
