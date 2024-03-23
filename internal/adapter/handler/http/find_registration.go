package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/employee-registration/internal/core/exception"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"net/http"
)

type FindAllRegisterDayHandler struct {
	useCase port.FindAllRegisterDayUseCase
}

func NewFindRegisterDayByUserIdHandler(useCase port.FindAllRegisterDayUseCase) *FindAllRegisterDayHandler {
	return &FindAllRegisterDayHandler{useCase: useCase}
}

func (h *FindAllRegisterDayHandler) Handle(c *gin.Context) {
	var requestURIParams dto.FindAllRegisterRequestURI

	err := c.ShouldBindUri(&requestURIParams)

	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid param user_id", err))
		return
	}

	c.JSON(http.StatusOK, h.useCase.FindAllRegisterDayByUserId(requestURIParams.UserId))
}
