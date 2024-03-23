package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/employee-registration/internal/core/exception"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"net/http"
)

type FindAllDailyRegistryHandler struct {
	useCase port.FindAllDailyRegistryUseCase
}

func NewFindAllDailyRegisterHandler(useCase port.FindAllDailyRegistryUseCase) *FindAllDailyRegistryHandler {
	return &FindAllDailyRegistryHandler{useCase: useCase}
}

func (h *FindAllDailyRegistryHandler) Handle(c *gin.Context) {
	var requestURIParams dto.FindDailyRegistryRequestURI

	err := c.ShouldBindUri(&requestURIParams)

	if err != nil {
		c.Error(exception.NewInvalidDataException("invalid param user_id", err))
		return
	}

	dailyRegisters, err := h.useCase.FindAllDailyRegistry(requestURIParams.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dailyRegisters)
}
