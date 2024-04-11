// handlers/epic.go
package handlers

import (
	"encoding/json"
	"net/http"
	"voter-search/utils"
)

func CommonStatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stateResponse, err := utils.FetchCommonState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stateResponse)
}
