package repository

import (
	"time"

	"database/sql"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

const query = `SELECT e.name 
		,e.position
     	,u.email
     	,r.date_time
	FROM register r
	INNER JOIN employee e ON e.id = r.employee_id
	INNER JOIN user u ON u.id = e.user_id
	WHERE u.id = ?
	  AND EXTRACT(MONTH FROM r.date_time) = ?
	  AND EXTRACT(YEAR FROM r.date_time) = ?
	ORDER BY r.date_time`

type mirrorRepository struct {
	db *sql.DB
}

func NewMirrorRepository(db *sql.DB) port.MirrorRepository {
	return mirrorRepository{db: db}
}

func (r mirrorRepository) GetMirror(userId, month, year int) (*domain.Mirror, error) {
	rows, err := r.db.Query(query, userId, month, year)
	if err != nil {
		return nil, err
	}

	mirror := &domain.Mirror{}

	for rows.Next() {
		register := time.Time{}
		err = rows.Scan(&mirror.Name, &mirror.Position, &mirror.Email, &register)
		if err != nil {
			return nil, err
		}
		mirror.Registers = append(mirror.Registers, register)
	}

	return mirror, nil
}
