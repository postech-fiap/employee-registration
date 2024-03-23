package repository

import (
	"database/sql"
	"github.com/postech-fiap/employee-registration/internal/core/dto"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

type registerDayRepository struct {
	db *sql.DB
}

func NewFindRegisterDayByUserIdRepository(db *sql.DB) port.FindAllRegisterDayRepository {
	return registerDayRepository{db: db}
}

func (r registerDayRepository) FindAllRegisterDayByUserId(userId uint64) []dto.RegisterDay {
	query := `SELECT e.name 
		,e.position
     	,r.date_time 
	FROM register r
	INNER JOIN employee e ON e.id = r.employee_id
	INNER JOIN user u ON u.id = e.user_id
	WHERE u.id = ?
	AND DAY(r.date_time) = DAY(CURDATE())`

	var registerDayByUser []dto.RegisterDay
	rows, err := r.db.Query(query, userId)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var name string
		var position string
		var dateTime string

		err = rows.Scan(&name, &position, &dateTime)
		if err != nil {
			panic(err.Error())
		}

		registerDayByUser = append(registerDayByUser, dto.RegisterDay{Name: name, Position: position, DateTime: dateTime})
	}

	return registerDayByUser
}
