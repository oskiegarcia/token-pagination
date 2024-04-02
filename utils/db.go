package utils

import "database/sql"

func QueryData(offset int, pageSize int) ([]*Data, int, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, 0, err
	}
	defer db.Close()

	// Query data from database
	rows, err := db.Query("SELECT id, name FROM items ORDER BY id LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Parse query results
	var data []*Data
	var totalCount int
	for rows.Next() {
		var item Data
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, 0, err
		}
		data = append(data, &item)
	}

	// Query total count
	err = db.QueryRow("SELECT COUNT(*) FROM items").Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return data, totalCount, nil
}
