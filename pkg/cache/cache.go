package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
)

func RecoverData(key string) (map[string]interface{}, error) {
	redis, err := RedisCnn()
	if err != nil || redis == nil {
		return nil, err
	}

	val, err := redis.Get(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SaveData(key string, data map[string]interface{}, expiration time.Duration) error {
	redis, err := RedisCnn()
	if err != nil || redis == nil {
		return err
	}

	if expiration == 0 {
		expiration = 10 * time.Minute
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = redis.Set(context.TODO(), key, jsonData, expiration).Err()
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
