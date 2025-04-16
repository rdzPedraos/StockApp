package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DB       dbConfig
	Server   serverConfig
	Services servicesConfig
}

var once sync.Once
var instance Config

func Get() *Config {
	once.Do(func() {
		instance = load()
	})
	return &instance
}

func load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DB:       getDBConfig(),
		Server:   getServerConfig(),
		Services: getServicesConfig(),
	}
}
