package main

import (
	"log"
	"net/http"

	"github.com/FrAigner/Message-Broker/apiEndpoints"
)

func main() {
	// TODO: Datenbankverbindung herstellen
	// TODO: Let's Encrypt-Zertifikat abrufen und aktualisieren

	// API-Endpunkte definieren
	http.HandleFunc("/api/nachrichten", apiEndpoints.HandleNachrichten)
	http.HandleFunc("/health", apiEndpoints.HandleHealthCheck)

	// Server starten
	log.Println("Server gestartet")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server konnte nicht gestartet werden: ", err)
	}
}
