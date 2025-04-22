package stockRating

import (
	"app/config"
	"app/lib/apiClient"
	"log"
)

type service struct {
	client apiClient.APIClient
}

func Service() *service {
	baseURL := config.Services.Stock.ApiUrl
	apiKey := config.Services.Stock.ApiKey

	if baseURL == "" || apiKey == "" {
		log.Fatal("API URL or API Key is not set")
	}

	apiClient := apiClient.New(baseURL)
	apiClient.SetAuthToken(apiKey)

	return &service{
		client: *apiClient,
	}
}
