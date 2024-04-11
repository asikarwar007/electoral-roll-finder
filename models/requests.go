// models/requests.go
package models

type EpicSearchRequest struct {
	IsPortal    bool   `json:"isPortal"`
	EpicNumber  string `json:"epicNumber"`
	StateCd     string `json:"stateCd"`
	CaptchaId   string `json:"captchaId"`
	CaptchaData string `json:"captchaData"`
	SecurityKey string `json:"securityKey"`
}
