package stockRating

import (
	"app/config"
	"app/lib/apiClient"
	"log"
)

type service struct {
	client apiClient.APIClient
}

var Service service

func init() {
	baseURL := config.Services.Stock.API_URL
	apiKey := config.Services.Stock.API_KEY

	if baseURL == "" || apiKey == "" {
		log.Fatal("API URL or API Key is not set")
	}

	apiClient := apiClient.New(baseURL)
	apiClient.SetAuthToken(apiKey)

	Service = service{
		client: *apiClient,
	}
}
