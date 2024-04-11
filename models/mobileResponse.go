// models/mobileResponse.go
package models

type MobileAPIResponse struct {
	File         *string `json:"file"`
	Message      string  `json:"message"`
	Payload      *string `json:"payload"`
	RefId        *string `json:"refId"`
	Status       *string `json:"status"`
	ShowNextFlow *bool   `json:"showNextFlow"`
	StatusCode   *int    `json:"statusCode"`
}
