package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/exception"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"net/http"
)

type registerService struct {
	registerUseCase port.RegisterUseCaseInterface
}

func NewRegisterService(registerUseCase port.RegisterUseCaseInterface) *registerService {
	return &registerService{
		registerUseCase: registerUseCase,
	}
}

func (r *registerService) Register(c *gin.Context) {
	var userId = c.GetHeader("user-id")
	err := c.ShouldBindHeader(&userId)

	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid param id", err))
		return
	}

	newRegister := domain.Register{
		ID: userId,
	}

	err = r.registerUseCase.PublishNewRegistry(newRegister)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusAccepted)

}
