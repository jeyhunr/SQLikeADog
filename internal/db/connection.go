package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeyhunr/SQLikeADog/internal/utils"
)

func Connect(config utils.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.Host, config.DBName)
	return sql.Open("mysql", dsn)
}
