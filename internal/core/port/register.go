package port

import "github.com/postech-fiap/employee-registration/internal/core/domain"

type RegisterUseCaseInterface interface {
	PublishNewRegistry(register domain.Register) error
	Insert(register *domain.Register) error
}

type RegisterQueuePublisherInterface interface {
	PublishRegistry(register domain.Register) error
}

type RegisterRepositoryInterface interface {
	Insert(register *domain.Register) error
}
