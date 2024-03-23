package repository

import (
	"database/sql"
	"github.com/postech-fiap/employee-registration/internal/core/domain"
	"github.com/postech-fiap/employee-registration/internal/core/domain/entity"
	"github.com/postech-fiap/employee-registration/internal/core/port"
	"time"
)

type registerRepository struct {
	db *sql.DB
}

func NewRegisterRepository(db *sql.DB) port.RegisterRepositoryInterface {
	return registerRepository{db: db}
}
func NewFindRegisterDayByUserIdRepository(db *sql.DB) port.FindAllDailyRegistryRepository {
	return registerRepository{db: db}
}

func (r registerRepository) FindAllDailyRegistry(userId uint64) (*entity.DailyRegistry, error) {
	query := `SELECT e.name 
		,e.position
     	,r.date_time 
	FROM register r
	INNER JOIN employee e ON e.id = r.employee_id
	INNER JOIN user u ON u.id = e.user_id
	WHERE u.id = ?
	AND DAY(r.date_time) = DAY(CURDATE())
	ORDER BY r.date_time`

	rows, err := r.db.Query(query, userId)

	if err != nil {
		return nil, err
	}

	registersDay := &entity.DailyRegistry{}

	for rows.Next() {
		register := time.Time{}
		err = rows.Scan(&registersDay.Name, &registersDay.Position, &register)
		if err != nil {
			return nil, err
		}
		registersDay.DailyRegistry = append(registersDay.DailyRegistry, register)
	}

	return registersDay, nil
}

func (r registerRepository) Insert(register *domain.Register) error {
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
