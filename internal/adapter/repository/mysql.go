package repository

import (
	"database/sql"
)

type mySQLRepository struct {
	client *sql.DB
}

//func NewMySQLRepository(client *sql.database) port.EmployeeRepositoryInterface {
//	return &mySQLRepository{
//		client: client,
//	}
//}

// methods
