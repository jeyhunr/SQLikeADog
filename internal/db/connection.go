package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect initializes the database connection
func Connect(host, port, user, password, dbName string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}
