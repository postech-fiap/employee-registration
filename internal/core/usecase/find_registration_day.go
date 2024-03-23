package usecase

import (
	"github.com/postech-fiap/employee-registration/internal/core/dto"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

type registerDayUseCase struct {
	repository port.FindAllRegisterDayRepository
}

func FindAllRegisterDayByUserIdUseCase(repository port.FindAllRegisterDayRepository) port.FindAllRegisterDayUseCase {
	return registerDayUseCase{repository: repository}
}

func (r registerDayUseCase) FindAllRegisterDayByUserId(userId uint64) []dto.RegisterDay {
	return r.repository.FindAllRegisterDayByUserId(userId)
}
