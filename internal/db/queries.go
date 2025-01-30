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

// GetTableData returns all data from a table
func GetTableData(dbName, tableName string) ([]string, [][]string, error) {
	if _, err := DB.Exec("USE " + dbName); err != nil {
		return nil, nil, fmt.Errorf("failed to use database: %v", err)
	}

	// Get columns
	rows, err := DB.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 1", tableName))
	if err != nil {
		return nil, nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}
	rows.Close()

	// Get data
	rows, err = DB.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var data [][]string
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, nil, err
		}

		row := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				row[i] = "NULL"
			} else {
				row[i] = fmt.Sprintf("%v", val)
			}
		}
		data = append(data, row)
	}

	return columns, data, nil
}
