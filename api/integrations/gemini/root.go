package gemini

import (
	"app/config"
	"app/lib/apiClient"
	"log"
)

type service struct {
	client *apiClient.APIClient
	apiKey string
}

func Service() *service {
	apiKey := config.Services.Gemini.ApiKey
	apiUrl := config.Services.Gemini.ApiUrl

	if apiKey == "" || apiUrl == "" {
		log.Fatal("Gemini API Key or API URL is not set")
	}

	apiClient := apiClient.New(apiUrl)

	return &service{
		client: apiClient,
		apiKey: apiKey,
	}
}
