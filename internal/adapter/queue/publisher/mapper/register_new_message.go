package mapper

import (
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/publisher/dto"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
)

func DomainToRegisterNewMessage(registerDomain domain.Register) *dto.RegisterNewMessage {
	return &dto.RegisterNewMessage{
		ID: registerDomain.ID,
	}
}
