package models

// CaptchaResponse matches the JSON structure of the external API response for captcha generation.
type CaptchaResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Captcha    string `json:"captcha"`
	ID         string `json:"id"`
}
