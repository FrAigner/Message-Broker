package appFunction

import (
	"database/sql"
	"fmt"

	"github.com/FrAigner/Message-Broker/config"
)

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
func ConnectToDatabase() (*sql.DB, error) {
	// Lade die Konfiguration aus der config.json-Datei
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Erstelle eine Datenbank-DSN-Zeichenfolge aus den Konfigurationsoptionen
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/app_db", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port)

	// Verbinde mit der Datenbank
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Überprüfe, ob die Verbindung funktioniert
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
