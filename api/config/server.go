package config

import (
	"app/lib/helper"
	"os"
)

type serverConfig struct {
	Host string
	Port string
	Url  string
}

var Server = serverConfig{
	Host: helper.Coalesce(os.Getenv("APP_HOST"), "localhost"),
	Port: helper.Coalesce(os.Getenv("APP_PORT"), "3000"),
	Url:  helper.Coalesce(os.Getenv("APP_URL"), "http://localhost:3000"),
}
