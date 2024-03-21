package repository

import (
	"database/sql"
)

type mySQLRepository struct {
	client *sql.DB
}

//func NewMySQLRepository(client *sql.DB) port.EmployeeRepositoryInterface {
//	return &mySQLRepository{
//		client: client,
//	}
//}

// methods
