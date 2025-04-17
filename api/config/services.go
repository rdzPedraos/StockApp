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
	Stock serviceConfig
}

var Services = servicesConfig{
	Stock: serviceConfig{
		ApiUrl: helper.Coalesce(os.Getenv("STOCK_API_URL"), ""),
		ApiKey: helper.Coalesce(os.Getenv("STOCK_API_KEY"), ""),
	},
}
