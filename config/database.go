package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	// Lade die Konfiguration aus der config.json-Datei
	config, err := LoadConfig()
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
