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
		registerParsed, _ := time.Parse("2006-01-02 15:4:5", register)
		if !before.IsZero() && before.Compare(registerParsed) == -1 && (i+1)%2 == 0 {
			hours += registerParsed.Sub(before)
		}
		before = registerParsed
	}
	register.Hours = hours.String()

	return register, nil
}
