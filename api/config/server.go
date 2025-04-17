package config

import (
	"app/lib/helper"
	"os"
)

type serverConfig struct {
	Host string
	Port string
}

var Server = serverConfig{
	Host: helper.Coalesce(os.Getenv("HOST"), "localhost"),
	Port: helper.Coalesce(os.Getenv("PORT"), "3000"),
}
