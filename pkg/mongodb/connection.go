package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
	"time"
)

var (
	clientInstance *mongo.Client
	once           sync.Once
)

func MongodbCnn() (*mongo.Client, error) {
	once.Do(func() {
		username := os.Getenv("MONGO_USER")
		password := os.Getenv("MONGO_PASS")
		host := os.Getenv("MONGO_SERVER")
		database := os.Getenv("MONGO_DB")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", username, password, host, database)
		clientOptions := options.Client().ApplyURI(uri)

		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstance = nil
		}

		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstance = nil
		}

		clientInstance = client
	})

	if clientInstance == nil {
		return nil, fmt.Errorf("failed to create a MongoDB client")
	}

	return clientInstance, nil
}
