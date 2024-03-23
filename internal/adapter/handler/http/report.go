package http

import (
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"net/http"
	"strconv"
)

type reportHandler struct {
	useCase port.ReportUseCase
}

func NewReportHandler(useCase port.ReportUseCase) reportHandler {
	return reportHandler{useCase: useCase}
}

func (h reportHandler) Handle(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetHeader("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.useCase.GenarateMirrorReport(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, "O relatório foi gerado com sucesso e estará disponível em breve no seu e-mail.")
}
