package usecase

import (
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type mirrorUseCase struct {
	repository port.MirrorRepository
}

func NewMirrorUseCase(repository port.MirrorRepository) port.MirrorUseCase {
	return mirrorUseCase{repository: repository}
}

func (r mirrorUseCase) GetMirror(userId int) (*domain.Mirror, error) {
	var before time.Time
	var hours time.Duration
	date := time.Now().AddDate(0, -1, 0)
	mirror, err := r.repository.GetMirror(userId, int(date.Month()), date.Year())
	if err != nil {
		return nil, err
	}

	mirror.Month = date.Month().String()
	mirror.Year = date.Year()

	for i, register := range mirror.Registers {
		if !before.IsZero() && before.Compare(register) == -1 && (i+1)%2 == 0 {
			hours += register.Sub(before)
		}
		before = register
	}
	mirror.Hours = hours.String()

	return mirror, nil
}
