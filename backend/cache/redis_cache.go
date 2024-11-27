package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Create a Redis client
var RedisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis server address (update as needed)
})

// SetCache stores the key-value pair in Redis with an expiration time
func SetCache(key, value string) error {
	// Set cache with expiration time of 1 hour
	return RedisClient.Set(ctx, key, value, 1*time.Hour).Err()
}

// GetCache retrieves the value from Redis by key
func GetCache(key string) (string, error) {
	// Get cached value
	return RedisClient.Get(ctx, key).Result()
}
