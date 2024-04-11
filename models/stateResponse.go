package models

// CaptchaResponse matches the JSON structure of the external API response for captcha generation.
type StateResponse struct {
	StateId        int     `json:"stateId"`
	CountryCd      string  `json:"countryCd"`
	StateCd        string  `json:"stateCd"`
	StateType      string  `json:"stateType"`
	StateName      string  `json:"stateName"`
	StateNameHindi string  `json:"stateNameHindi"`
	EffectiveFrom  string  `json:"effectiveFrom"`
	EffectiveTo    string  `json:"effectiveTo"`
	IsActive       string  `json:"isActive"`
	CreatedBy      string  `json:"createdBy"`
	CreatedDttm    string  `json:"createdDttm"`
	ModifiedBy     *string `json:"modifiedBy,omitempty"`
	ModifiedDttm   *string `json:"modifiedDttm,omitempty"`
}
