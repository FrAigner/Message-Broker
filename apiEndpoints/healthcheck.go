package apiEndpoints

import (
	"encoding/json"
	"net/http"
	"time"

	appFunction "github.com/FrAigner/Message-Broker/functions/appFunction"
)

type HealthCheck struct {
	Healthiness string        `json:"healthiness"`
	Update      time.Time     `json:"update"`
	Hints       []interface{} `json:"hints"`
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	// Check configuration
/* 	funct, err := config.LoadConfig()
		if err != nil {
	   		w.WriteHeader(http.StatusInternalServerError)
	   		w.Write([]byte(`{"healthiness": "red", "hints": [{"config": "failed to load config"}]}`))
	   		return
	   	} */
	   
	// Check database connection
	db, err := appFunction.ConnectToDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"healthiness": "red", "hints": [{"database": "failed to connect to database"}]}`))
		return
	}
	defer db.Close()

	// Check database table
	err = appFunction.CreateTableIfNotExists(db, "messages")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"healthiness": "red", "hints": [{"database": "failed to create table"}]}`))
		return
	}

	// Build health check response
	now := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	hints := []map[string]string{
		{"config": "ok"},
		{"database": "ok"},
	}
	response := map[string]interface{}{
		"healthiness": "green",
		"update":      now,
		"hints":       hints,
	}

	// Send response
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"healthiness": "red", "hints": [{"config": "failed to create response"}]}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
