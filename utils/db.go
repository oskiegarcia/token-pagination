package utils

import (
	"database/sql"
	"fmt"
	"log"
)

// DBLogger wraps *sql.DB and logs queries before execution
type DBLogger struct {
	*sql.DB
}

// Query executes a query and logs the SQL statement before execution
func (db *DBLogger) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf("Executing query: %s", query)
	return db.DB.Query(query, args...)
}

// Exec executes a query and logs the SQL statement before execution
func (db *DBLogger) Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Executing query: %s", query)
	return db.DB.Exec(query, args...)
}

// QueryData is function to query the data
func QueryData(lastID int, pageSize int) ([]*Data, int, int) {
	// Open database connection
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, -1, 0
	}
	defer db.Close()

	// Wrap *sql.DB with DBLogger
	dbWrapper := &DBLogger{db}

	// Query data from database
	rows, err := dbWrapper.Query("SELECT id, name FROM items WHERE id > ? ORDER BY id LIMIT ?", lastID, pageSize)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, -1, 0
	}
	defer rows.Close()

	// Parse query results
	var data []*Data
	var nextID int
	var totalCount int
	for rows.Next() {
		var item Data
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, -1, 0
		}
		data = append(data, &item)
		nextID = item.ID
	}

	// Query total count
	err = db.QueryRow("SELECT COUNT(*) FROM items").Scan(&totalCount)
	if err != nil {
		fmt.Println("Error querying database for total count:", err)
		return nil, -1, 0
	}

	return data, nextID, totalCount
}
