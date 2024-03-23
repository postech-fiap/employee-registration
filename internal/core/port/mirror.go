package port

import (
	"github.com/postech-fiap/employee-registration/internal/core/domain"
)

type MirrorRepository interface {
	GetMirror(userId, month, year int) (*domain.Mirror, error)
}

type MirrorUseCase interface {
	GetMirror(userId int) (*domain.Mirror, error)
}
