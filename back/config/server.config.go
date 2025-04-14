package config

import (
	"app/utils"
	"os"
)

type serverConfig struct {
	Host string
	Port string
}

func getServerConfig() serverConfig {
	return serverConfig{
		Host: utils.Coalesce(os.Getenv("HOST"), "localhost"),
		Port: utils.Coalesce(os.Getenv("PORT"), "3000"),
	}
}
