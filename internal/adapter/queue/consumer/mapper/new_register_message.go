package mapper

import (
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/consumer/dto"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
)

func MapNewRegisterMessageToDomain(registerDTO *dto.NewRegisterMessage) *domain.Register {
	return &domain.Register{
		ID: registerDTO.ID,
	}
}
