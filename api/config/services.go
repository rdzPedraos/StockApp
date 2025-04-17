package config

import (
	"app/lib/helper"
	"os"
)

type serviceConfig struct {
	API_URL string
	API_KEY string
}

type servicesConfig struct {
	Stock serviceConfig
}

var Services = servicesConfig{
	Stock: serviceConfig{
		API_URL: helper.Coalesce(os.Getenv("STOCK_API_URL"), ""),
		API_KEY: helper.Coalesce(os.Getenv("STOCK_API_KEY"), ""),
	},
}
