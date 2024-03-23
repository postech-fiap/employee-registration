package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/postech-fiap/employee-registration/cmd/config"
)

var connection *sql.DB = nil

func OpenConnection(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Schema))
	if err != nil {
		return nil, err
	}

	connection = db

	err = testConnection()
	if err != nil {
		return nil, err
	}

	return connection, nil

}

func CloseConnection() {
	if err := connection.Close(); err != nil {
		panic(err)
	}
	fmt.Println("MySQL disconnected")
}

func testConnection() error {
	_, err := connection.Query("select now()")
	if err != nil {
		return err
	}
	return nil
}
