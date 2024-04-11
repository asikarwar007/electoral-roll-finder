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
type MobileSendOTPRequest struct {
	MobNo       string `json:"mobNo"`
	StateCd     string `json:"stateCd"`
	CaptchaId   string `json:"captchaId"`
	CaptchaData string `json:"captchaData"`
	SecurityKey string `json:"securityKey"`
}
type MobileSearchRequest struct {
	MobileNumber string `json:"mobileNumber"`
	StateCd      string `json:"stateCd"`
	Otp          string `json:"otp"`
}
