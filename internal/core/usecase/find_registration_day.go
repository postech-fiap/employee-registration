package usecase

import (
	"github.com/postech-fiap/employee-registration/internal/core/domain/entity"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type registerDayUseCase struct {
	repository port.FindAllDailyRegistryRepository
}

func FindAllRegisterDayByUserIdUseCase(repository port.FindAllDailyRegistryRepository) port.FindAllDailyRegistryUseCase {
	return registerDayUseCase{repository: repository}
}

func (r registerDayUseCase) FindAllDailyRegistry(userId uint64) (*entity.DailyRegistry, error) {
	var before time.Time
	var hours time.Duration

	register, err := r.repository.FindAllDailyRegistry(userId)
	if err != nil {
		return nil, err
	}

	for i, register := range register.DailyRegistry {
		if !before.IsZero() && before.Compare(register) == -1 && (i+1)%2 == 0 {
			hours += register.Sub(before)
		}
		before = register
	}
	register.Hours = hours.String()

	return register, nil
}
