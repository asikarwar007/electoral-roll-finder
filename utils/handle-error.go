package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"voter-search/models"
)

// sendJSONError sends a JSON formatted error message to the client.
func SendJSONError(w http.ResponseWriter, errMsg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error:   true,
		Message: errMsg,
	})
}

// handleHTTPError handles non-200 HTTP responses and constructs a detailed error message.
func HandleHTTPError(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// If we cannot read the body, use a generic error message
		return fmt.Errorf("HTTP %d: unable to read response body", resp.StatusCode)
	}

	// Attempt to parse the body as the initial error response structure
	var initialErrResp struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(bodyBytes, &initialErrResp); err != nil {
		return fmt.Errorf("%s", string(bodyBytes)) // Fall back to raw message
	}

	// Check if message contains JSON
	if len(initialErrResp.Message) > 0 && initialErrResp.Message[0] == '{' {
		var detailedErr models.DetailedErrorResponse
		if err := json.Unmarshal([]byte(initialErrResp.Message), &detailedErr); err == nil {
			return fmt.Errorf("%s", detailedErr.Message)
		}
	}

	// Use the message from the initial error response if the nested JSON is not valid or not present
	return fmt.Errorf("%s", initialErrResp.Message)

	// errMsg := string(bodyBytes) // Use the body as the error message if readable
	// switch resp.StatusCode {
	// case http.StatusBadRequest:
	// 	return fmt.Errorf("bad request to API: %s", errMsg)
	// case http.StatusUnauthorized:
	// 	return fmt.Errorf("unauthorized access to external API: %s", errMsg)
	// case http.StatusForbidden:
	// 	return fmt.Errorf("forbidden access to external API: %s", errMsg)
	// case http.StatusNotFound:
	// 	return fmt.Errorf("API endpoint not found: %s", errMsg)
	// default:
	// 	return fmt.Errorf("external API returned HTTP %d: %s", resp.StatusCode, errMsg)
	// }
}
