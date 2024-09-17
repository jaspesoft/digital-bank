package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var (
	clientInstance *redis.Client
	once           sync.Once
)

func RedisCnn() (*redis.Client, error) {
	once.Do(func() {
		db := 0

		if os.Getenv("GO_ENV") != "dev" {
			db = 1
		}

		uri := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"))
		// Create a Redis client
		client := redis.NewClient(&redis.Options{
			Addr:     uri,
			Password: "",
			DB:       db, // Use default DB
		})

		// Create a context
		ctx := context.Background()

		// Ping the Redis server and check if any errors occurred
		_, err := client.Ping(ctx).Result()
		if err != nil {
			clientInstance = nil
			fmt.Println("failed to create a Redis client")
		}

		clientInstance = client
	})

	if clientInstance == nil {
		return nil, fmt.Errorf("failed to create a Redis client")
	}

	return clientInstance, nil
}
