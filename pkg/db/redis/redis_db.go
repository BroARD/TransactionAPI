package redis

import (
	"TransactionAPI/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := ":6379"

	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	return client
}