// handlers/epic.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"voter-search/models"
)

func VoterDetailsSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.EpicSearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the received JSON
	log.Printf("Received EPIC search request: %+v\n", req)

	response := models.EpicSearchResponse{
		Error:   false,
		Message: "got param",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
