// models/responses.go
package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DetailedErrorResponse struct {
	Status       string      `json:"status"`
	StatusCode   int         `json:"statusCode"`
	Message      interface{} `json:"message"` // Use interface{} to accept any data type
	ShowNextFlow bool        `json:"showNextFlow"`
	RefId        interface{} `json:"refId"`
	Payload      interface{} `json:"payload"`
	File         interface{} `json:"file"`
}

// CustomTime type to handle non-standard time formats in JSON
type CustomTime struct {
	time.Time
}

// UnmarshalJSON for CustomTime to parse time in the "2006-01-02 15:04:05" format
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	const layout = "2006-01-02 15:04:05"
	s := string(b)
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	s = s[1 : len(s)-1] // Strip the quotes
	t, err := time.Parse(layout, s)
	if err != nil {
		return fmt.Errorf("parsing time %q as %q: %w", s, layout, err)
	}
	ct.Time = t
	return nil
}

// TrimmedString is a string that trims whitespace around it when unmarshaled from JSON.
type TrimmedString string

// UnmarshalJSON implements the json.Unmarshaler interface for TrimmedString.
func (ts *TrimmedString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*ts = TrimmedString(strings.TrimSpace(s))
	return nil
}

// ErrorResponse represents a JSON structure for API error responses.
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
