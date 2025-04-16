package config

import (
	"app/utils"
	"os"
)

type serviceConfig struct {
	API_URL string
	API_KEY string
}

type servicesConfig struct {
	Stock serviceConfig
}

func getServicesConfig() servicesConfig {
	return servicesConfig{
		Stock: serviceConfig{
			API_URL: utils.Coalesce(os.Getenv("STOCK_API_URL"), ""),
			API_KEY: utils.Coalesce(os.Getenv("STOCK_API_KEY"), ""),
		},
	}
}
