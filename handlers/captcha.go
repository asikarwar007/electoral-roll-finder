// handlers/epic.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"voter-search/utils"
)

func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	captchaResponse, err := utils.FetchGenerateCaptcha()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Check if the 'render' query parameter is present
	render := r.URL.Query().Get("render") == "true"
	if render {
		// Respond with HTML
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<img src="data:image/jpeg;base64,%s" alt="Captcha" /><div>Captcha ID: %s</div>`, captchaResponse.Captcha, captchaResponse.ID)
	} else {
		// Respond with JSON
		response := map[string]interface{}{
			"captcha": captchaResponse.Captcha,
			"id":      captchaResponse.ID,
			"error":   false,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
