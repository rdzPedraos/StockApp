package financial

import (
	"app/config"
	"app/lib/apiClient"
	"log"
)

type service struct {
	client apiClient.APIClient
	apiKey string
}

func Service() *service {
	baseURL := config.Services.Financial.ApiUrl
	apiKey := config.Services.Financial.ApiKey

	if baseURL == "" || apiKey == "" {
		log.Fatal("API URL or API Key is not set")
	}

	apiClient := apiClient.New(baseURL)

	return &service{
		client: *apiClient,
		apiKey: apiKey,
	}
}
