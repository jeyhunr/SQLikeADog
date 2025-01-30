package db

import (
	"fmt"
)

// ListDatabases returns a list of all databases
func ListDatabases() ([]string, error) {
	rows, err := DB.Query("SHOW DATABASES")
	if err != nil {
		return nil, fmt.Errorf("failed to list databases: %v", err)
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return nil, fmt.Errorf("failed to scan database name: %v", err)
		}
		databases = append(databases, dbName)
	}
	return databases, nil
}

// ListTables returns a list of all tables in a database
func ListTables(dbName string) ([]string, error) {
	if _, err := DB.Exec("USE " + dbName); err != nil {
		return nil, fmt.Errorf("failed to use database: %v", err)
	}

	rows, err := DB.Query("SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %v", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %v", err)
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}
