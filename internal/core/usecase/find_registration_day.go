package usecase

import (
	"github.com/postech-fiap/employee-registration/internal/core/dto"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type registerDayUseCase struct {
	repository port.FindAllRegisterDayRepository
}

func FindAllRegisterDayByUserIdUseCase(repository port.FindAllRegisterDayRepository) port.FindAllRegisterDayUseCase {
	return registerDayUseCase{repository: repository}
}

func (r registerDayUseCase) FindAllRegisterDayByUserId(userId uint64) (*dto.RegisterDay, error) {
	var before time.Time
	var hours time.Duration

	register, err := r.repository.FindAllRegisterDayByUserId(userId)

	if err != nil {
		return nil, err
	}

	for i, register := range register.Registers {
		registerParsed, _ := time.Parse("2006-01-02 15:4:5", register)
		if !before.IsZero() && before.Compare(registerParsed) == -1 && (i+1)%2 == 0 {
			hours += registerParsed.Sub(before)
		}
		before = registerParsed
	}
	register.Hours = hours.String()

	return register, nil
}
