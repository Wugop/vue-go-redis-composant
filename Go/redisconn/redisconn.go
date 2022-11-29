package redisconn

import (
	"github.com/go-redis/redis"
)

func GetRedisConnection() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
