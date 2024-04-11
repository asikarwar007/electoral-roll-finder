package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"voter-search/models" // Adjust the import path according to your project structure
)

func FetchEpicInfo(requestData models.EpicSearchRequest) ([]models.EpicAPIResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s/api/v1/elastic/search-by-epic-from-national-display", apiBaseURL)

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request data: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error making request to external API: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			errMsg = string(bodyBytes) // Use the body as error message if readable
		} else {
			errMsg = "unable to read response body"
		}

		if resp.StatusCode == http.StatusBadRequest {
			return nil, fmt.Errorf("bad request to external API: %s", errMsg)
		}

		// Handle other status codes as needed
		return nil, fmt.Errorf("external API returned HTTP %d: %s", resp.StatusCode, errMsg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var responses []models.EpicAPIResponse
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return responses, nil
}

// FetchGenerateCaptcha calls the external captcha service and returns the response.
func FetchGenerateCaptcha() (*models.CaptchaResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	apiUrl := fmt.Sprintf("%s/api/v1/captcha-service/generateCaptcha", apiBaseURL)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to request captcha: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body) // Best effort to read the body for additional context
		return nil, fmt.Errorf("captcha API returned non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var captchaResponse models.CaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&captchaResponse); err != nil {
		return nil, fmt.Errorf("failed to decode captcha response: %w", err)
	}

	return &captchaResponse, nil
}

// FetchGenerateCaptcha calls the external captcha service and returns the response.
func FetchCommonState() ([]models.StateResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	apiUrl := fmt.Sprintf("%s/api/v1/common/states", apiBaseURL)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to request captcha: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body) // Best effort to read the body for additional context
		return nil, fmt.Errorf("captcha API returned non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var stateResponse []models.StateResponse

	if err := json.Unmarshal(body, &stateResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return stateResponse, nil
}
