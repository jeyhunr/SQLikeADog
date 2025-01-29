package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
