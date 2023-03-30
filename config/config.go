package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Config struct {
	MySQL       MySQLConfig       `json:"mysql"`
	API         APIConfig         `json:"api"`
	LetsEncrypt LetsEncryptConfig `json:"letsencrypt"`
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type APIConfig struct {
	APIKey         string          `json:"api_key"`
	MessageTargets []MessageTarget `json:"MessageTargets"`
}

type MessageTarget struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Webhook string `json:"webhook"`
}

type LetsEncryptConfig struct {
	Email   string   `json:"email"`
	Domains []string `json:"domains"`
}

func LoadConfig() (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configFile, err := os.Open(dir + "/config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config Config) error {

	if len(config.API.MessageTargets) == 0 {
		return errors.New("invalid config: missing message targets")
	}
	if config.MySQL.Host == "" {
		return errors.New("invalid config: missing DB server")
	}
	if config.MySQL.Port == 0 {
		return errors.New("invalid config: missing DB port")
	}
	if config.MySQL.Username == "" {
		return errors.New("invalid config: missing DB user")
	}
	if config.MySQL.Password == "" {
		return errors.New("invalid config: missing DB password")
	}

	if config.API.APIKey == "" {
		return errors.New("invalid config: missing API key")
	}
	if config.LetsEncrypt.Email == "" {
		return errors.New("invalid config: missing email for Let's Encrypt")
	}
	if len(config.LetsEncrypt.Domains) == 0 {
		return errors.New("invalid config: missing domains for Let's Encrypt")
	}
	return nil
}
