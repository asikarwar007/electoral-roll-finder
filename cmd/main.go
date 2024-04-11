// cmd/main.go
package main

import (
	"log"
	"net/http"
	"voter-search/handlers"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	http.HandleFunc("/epic", handlers.EpicSearchHandler)
	http.HandleFunc("/voter-details", handlers.VoterDetailsSearchHandler)
	http.HandleFunc("/mobile-send-otp", handlers.MobileSendOtpHandler)
	http.HandleFunc("/mobile-search", handlers.MobileSearchHandler)
	http.HandleFunc("/generate-captcha", handlers.GenerateCaptchaHandler)
	http.HandleFunc("/states", handlers.CommonStatesHandler)

	log.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
