package port

import "github.com/postech-fiap/employee-registration/internal/core/dto"

type FindAllRegisterDayRepository interface {
	FindAllRegisterDayByUserId(userId uint64) (*dto.RegisterDay, error)
}

type FindAllRegisterDayUseCase interface {
	FindAllRegisterDayByUserId(userId uint64) (*dto.RegisterDay, error)
}
