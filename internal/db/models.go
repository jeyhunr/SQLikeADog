package db

// Table represents a database table structure
type Table struct {
	Name    string
	Columns []Column
}

// Column represents a table column structure
type Column struct {
	Name     string
	Type     string
	Nullable bool
	Key      string
	Default  interface{}
	Extra    string
}

// TableData represents the data within a table
type TableData struct {
	Columns []string
	Rows    [][]interface{}
}
