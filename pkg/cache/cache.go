package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
)

func RecoverData(key string) (string, error) {
	redis, err := RedisCnn()
	if err != nil || redis == nil {
		return "", err
	}

	val, err := redis.Get(context.TODO(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil

}

func SaveData(key string, data string, expiration time.Duration) error {
	redis, err := RedisCnn()
	if err != nil || redis == nil {
		return err
	}

	if expiration == 0 {
		expiration = 10 * time.Minute
	}

	err = redis.Set(context.TODO(), key, data, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// InCache checks if the payload is in the cache.
func InCache(payload interface{}) bool {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error converting payload to JSON: %v", err)
	}

	// Create a buffer from the byte slice
	buffer := bytes.NewBuffer(data)

	redis, err := RedisCnn()
	if err != nil || redis == nil {
		return false
	}
	val, err := redis.Get(context.TODO(), buffer.String()).Result()

	if err != nil {
		if err.Error() != "redis: nil" {
			log.Fatalf("Error getting key: %v", err)
			return false
		}
	}

	if val == "" {
		timeInCache := 60 * time.Minute
		if os.Getenv("GO_ENV") == "dev" {
			timeInCache = 60 * time.Second
		}

		err = redis.Set(context.TODO(), buffer.String(), buffer.String(), timeInCache).Err()
		if err != nil {
			log.Fatalf("Error setting key: %v", err)
		}

		return false
	}

	return true
}
