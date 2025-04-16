package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func init() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	})
}
