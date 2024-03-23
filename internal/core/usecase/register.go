package usecase

import (
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/exception"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

type registerUseCase struct {
	registerRepository     port.RegisterRepositoryInterface
	registerQueuePublisher port.RegisterQueuePublisherInterface
}

func (r registerUseCase) PublishNewRegistry(register domain.Register) error {

	err := r.registerQueuePublisher.PublishRegistry(register)
	if err != nil {
		return exception.NewFailedDependencyException(err)

	}
	return nil
}

func (r registerUseCase) Insert(register *domain.Register) error {

	err := r.registerRepository.Insert(register)
	if err != nil {
		return exception.NewFailedDependencyException(err)
	}
	return nil
}

func NewRegisterUseCase(registerRepository port.RegisterRepositoryInterface, registerQueuePublisher port.RegisterQueuePublisherInterface) port.RegisterUseCaseInterface {
	return &registerUseCase{
		registerRepository:     registerRepository,
		registerQueuePublisher: registerQueuePublisher,
	}
}
