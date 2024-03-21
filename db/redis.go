// db/redis.go
package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// InitRedis initializes Redis connection
func InitRedis(addr string) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to ping Redis:", err)
	}
	redisClient = client
}

// GetRedisClient returns Redis client
func GetRedisClient() *redis.Client {
	return redisClient
}
