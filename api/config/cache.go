package config

import (
	"app/lib/helper"
	"os"
)

type cacheConfig struct {
	Address  string
	Password string
	DB       int
}

var Cache = cacheConfig{
	Address:  helper.Coalesce(os.Getenv("CACHE_ADDRESS"), "localhost:6379"),
	Password: helper.Coalesce(os.Getenv("CACHE_PASSWORD"), ""),
	DB:       0,
}
