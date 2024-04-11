// handlers/epic.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"voter-search/models"
	"voter-search/utils"
)

func EpicSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.EpicSearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	// Log the received JSON
	log.Printf("Received EPIC search request: %+v\n", req)

	apiResponse, err := utils.FetchEpicInfo(req)
	if err != nil {
		utils.SendJSONError(w, "Error fetching EPIC info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse[0].Content) // Respond with the "content" part of the API response
}
