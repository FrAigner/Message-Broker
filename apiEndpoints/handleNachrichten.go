package apiEndpoints

import (
	"fmt"
	"net/http"
)

func HandleNachrichten(w http.ResponseWriter, r *http.Request) {
	// TODO: Authentifizierung überprüfen
	// TODO: UUID für Anfrage generieren
	// TODO: Nachricht in Datenbank speichern
	// TODO: Nachricht an Ziele weiterleiten

	// Antwort zurückgeben
	fmt.Fprintf(w, "Nachricht empfangen und verarbeitet")
}
