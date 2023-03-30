package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/FrAigner/Message-Broker/endpoints"
)

func main() {
	// TODO: Datenbankverbindung herstellen
	// TODO: Let's Encrypt-Zertifikat abrufen und aktualisieren

	// API-Endpunkte definieren
	http.HandleFunc("/api/nachrichten", endpoints.HandleNachrichten)
	http.HandleFunc("/health", endpoints.HandleHealthCheck)
	// Server starten
	log.Println("Server gestartet")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server konnte nicht gestartet werden: ", err)
	}
}

type Config struct {
	MessageTargets map[string]string `json:"messageTargets"`
	DBServer       string            `json:"dbServer"`
	DBPort         int               `json:"dbPort"`
	DBUser         string            `json:"dbUser"`
	DBPassword     string            `json:"dbPassword"`
}

func LoadConfig() (*Config, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
