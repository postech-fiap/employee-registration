package port

import (
	"github.com/postech-fiap/employee-registration/internal/core/domain/entity"
)

type FindAllDailyRegistryRepository interface {
	FindAllDailyRegistry(userId uint64) (*entity.DailyRegistry, error)
}

type FindAllDailyRegistryUseCase interface {
	FindAllDailyRegistry(userId uint64) (*entity.DailyRegistry, error)
}
