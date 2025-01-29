package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"dbname"`
}

func saveConfig(config DBConfig, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func loadConfig(filename string) (DBConfig, error) {
	var config DBConfig
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func listDatabases(config DBConfig) ([]string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", config.User, config.Password, config.Host)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return nil, err
		}
		databases = append(databases, dbName)
	}
	return databases, nil
}

func connectToDatabase(user, password, host, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func executeQuery(query string) (string, error) {
	// Connect to the database
	db, err := connectToDatabase("user", "password", "localhost", "dbname")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Process the results
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	results := ""
	for rows.Next() {
		columnsData := make([]interface{}, len(columns))
		columnsPointers := make([]interface{}, len(columns))
		for i := range columnsData {
			columnsPointers[i] = &columnsData[i]
		}

		if err := rows.Scan(columnsPointers...); err != nil {
			return "", err
		}

		for i, colName := range columns {
			results += fmt.Sprintf("%s: %v\n", colName, columnsData[i])
		}
		results += "\n"
	}

	return results, nil
}
