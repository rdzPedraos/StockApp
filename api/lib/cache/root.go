package cache

import (
	"app/config"
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
	ctx context.Context
}

var cache Cache

func init() {
	cache = Cache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     config.Cache.Address,
			Password: config.Cache.Password,
			DB:       config.Cache.DB,
		}),
		ctx: context.Background(),
	}
}
