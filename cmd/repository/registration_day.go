package repository

import (
	"database/sql"
	"github.com/postech-fiap/employee-registration/internal/core/dto"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type registerDayRepository struct {
	db *sql.DB
}

func NewFindRegisterDayByUserIdRepository(db *sql.DB) port.FindAllRegisterDayRepository {
	return registerDayRepository{db: db}
}

func (r registerDayRepository) FindAllRegisterDayByUserId(userId uint64) (*dto.RegisterDay, error) {
	query := `SELECT e.name 
		,e.position
     	,r.date_time 
	FROM register r
	INNER JOIN employee e ON e.id = r.employee_id
	INNER JOIN user u ON u.id = e.user_id
	WHERE u.id = ?
	AND DAY(r.date_time) = DAY(CURDATE())`

	rows, err := r.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	registersDay := &dto.RegisterDay{}

	for rows.Next() {
		register := time.Time{}.String()
		err = rows.Scan(&registersDay.Name, &registersDay.Position, &register)
		if err != nil {
			return nil, err
		}
		registersDay.Registers = append(registersDay.Registers, register)
	}

	return registersDay, nil
}
