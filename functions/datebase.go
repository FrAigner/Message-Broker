package functions

import "database/sql"

func CreateTableIfNotExists(db *sql.DB, tableName string) error {
	// Check if the table exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = ?", tableName).Scan(&count)
	if err != nil {
		return err
	}

	// If the table doesn't exist, create it
	if count == 0 {
		_, err = db.Exec("CREATE TABLE " + tableName + " (id INT PRIMARY KEY, name VARCHAR(255))")
		if err != nil {
			return err
		}
	}

	return nil
}
