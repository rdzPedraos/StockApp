package config

import (
	"app/lib/helper"
	"os"
)

type serviceConfig struct {
	ApiUrl string
	ApiKey string
}

type servicesConfig struct {
	Stock     serviceConfig
	Financial serviceConfig
}

var Services = servicesConfig{
	Stock: serviceConfig{
		ApiUrl: helper.Coalesce(os.Getenv("STOCK_API_URL"), ""),
		ApiKey: helper.Coalesce(os.Getenv("STOCK_API_KEY"), ""),
	},
	Financial: serviceConfig{
		ApiUrl: helper.Coalesce(os.Getenv("FINANCIAL_API_URL"), ""),
		ApiKey: helper.Coalesce(os.Getenv("FINANCIAL_API_KEY"), ""),
	},
}
