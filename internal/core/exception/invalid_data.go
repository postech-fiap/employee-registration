package exception

import (
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"net/http"
)

type invalidDataException struct {
	baseException
}

func NewInvalidDataException(message string, error error) port.CustomExceptionInterface {
	return &invalidDataException{
		baseException{
			statusCode: http.StatusBadRequest,
			message:    message,
			error:      error,
		},
	}
}
