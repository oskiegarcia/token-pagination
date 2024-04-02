package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"token-pagination/utils"
)

func main() {
	// Open database connection
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Wrap *sql.DB with DBLogger
	dbWrapper := &utils.DBLogger{DB: db}

	// Create items table if not exists
	_, err = dbWrapper.Exec(`CREATE TABLE IF NOT EXISTS items (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert initial data
	initialData := []*utils.Data{
		{Name: "Item 1"},
		{Name: "Item 2"},
		{Name: "Item 3"},
		// Add more initial data as needed
	}

	for _, item := range initialData {
		_, err := db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
		if err != nil {
			fmt.Println("Error inserting data:", err)
			return
		}
	}

	fmt.Println("Initial data inserted successfully.")
}
