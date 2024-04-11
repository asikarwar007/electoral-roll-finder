// handlers/mobile.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"voter-search/models"
	"voter-search/utils"
)

func MobileSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.MobileSearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	// Log the received JSON
	log.Printf("Received EPIC search request: %+v\n", req)

	apiResponse, err := utils.FetchMobileSendOtp(req)
	if err != nil {
		utils.SendJSONError(w, "Error Sending Mobile OTP: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}
