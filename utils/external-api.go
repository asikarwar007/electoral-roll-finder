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
		return nil, HandleHTTPError(resp)
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

func FetchVoterSearchInfo(requestData models.VoterSearchRequest) ([]models.EpicAPIResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s/api/v1/elastic/search-by-details-from-state-display", apiBaseURL)

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
		return nil, HandleHTTPError(resp)
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

func FetchMobileSendOtp(requestData models.MobileSendOTPRequest) (*models.MobileAPIResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s/api/v1/elastic-otp/send-otp-search", apiBaseURL)

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
		return nil, HandleHTTPError(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var responses *models.MobileAPIResponse
	if err := json.Unmarshal(body, &responses); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return responses, nil
}

func FetchMobileDetails(requestData models.MobileSearchRequest) ([]models.EpicAPIResponse, error) {
	apiBaseURL := os.Getenv("API_URL")
	url := fmt.Sprintf("%s/api/v1/elastic/search-by-mobile-from-state-search-display", apiBaseURL)

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
		return nil, HandleHTTPError(resp)
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

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return nil, HandleHTTPError(resp)
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

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return nil, HandleHTTPError(resp)
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
