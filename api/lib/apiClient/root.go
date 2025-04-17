package apiClient

import (
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

func New(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
