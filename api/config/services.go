package config

import (
	"os"
)

type serviceConfig struct {
	ApiUrl string
	ApiKey string
}

type servicesConfig struct {
	Stock     serviceConfig
	Financial serviceConfig
	Gemini    serviceConfig
}

var Services = servicesConfig{
	Stock: serviceConfig{
		ApiUrl: os.Getenv("STOCK_API_URL"),
		ApiKey: os.Getenv("STOCK_API_KEY"),
	},
	Financial: serviceConfig{
		ApiUrl: os.Getenv("FINANCIAL_API_URL"),
		ApiKey: os.Getenv("FINANCIAL_API_KEY"),
	},
	Gemini: serviceConfig{
		ApiKey: os.Getenv("GEMINI_API_KEY"),
		ApiUrl: os.Getenv("GEMINI_API_URL"),
	},
}
