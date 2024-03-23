package repository

import (
	"database/sql"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/port"
)

type reportRepository struct {
	db *sql.DB
}

func NewRegisterRepository(db *sql.DB) port.RegisterRepositoryInterface {
	return reportRepository{db: db}
}

func (r reportRepository) Insert(register *domain.Register) error {
	query := "INSERT INTO register(date_time, employee_id)" +
		"SELECT CURRENT_TIMESTAMP, id from employee em where em.user_id = ?"

	_, err := r.db.Exec(query, register.ID)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		return err
	}

	return nil
}
