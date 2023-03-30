package main

import (
	"log"
	"net/http"

	"github.com/FrAigner/Message-Broker/apiEndpoints"
	"github.com/FrAigner/Message-Broker/config"
	"github.com/FrAigner/Message-Broker/functions"
)

func main() {
	// TODO: Datenbankverbindung herstellen
	// TODO: Let's Encrypt-Zertifikat abrufen und aktualisieren

	// API-Endpunkte definieren
	http.HandleFunc("/api/nachrichten", apiEndpoints.HandleNachrichten)
	http.HandleFunc("/health", apiEndpoints.HandleHealthCheck)

	// Server starten
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatal("Datenbankverbindung konnte nicht hergestellt werden: ", err)
	}

	err = functions.CreateTableIfNotExists(db, "apikey")
	if err != nil {
		log.Fatal("Tabelle konnte nicht erstellt werden: ", err)
	}

	log.Println("Server gestartet")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server konnte nicht gestartet werden: ", err)
	}
}
