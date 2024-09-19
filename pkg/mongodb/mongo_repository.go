package mongo

import (
	"context"
	"digital-bank/internal/system/domain/criteria"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type (
	Repository struct {
		CollectionName string
		MongoConvert   *MongoConverter
		MongoClient    *mongo.Client
	}
)

func NewMongoRepository(CollectionName string) *Repository {
	client, err := MongodbCnn()
	if err != nil {
		panic(err)
	}

	return &Repository{
		CollectionName: CollectionName,
		MongoClient:    client,
		MongoConvert:   NewMongoConverter(),
	}
}

func (repository *Repository) GetCollection() *mongo.Collection {
	return repository.MongoClient.Database(os.Getenv("MONGO_DB")).Collection(repository.CollectionName)
}

func (repository *Repository) SearchByCriteria(criteria *criteria.Criteria) (*mongo.Cursor, int, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := repository.MongoConvert.PrepareSearch(criteria)

	collection := repository.GetCollection()

	findOptions := options.Find()
	findOptions.SetSkip(int64(query.Skip))
	findOptions.SetLimit(int64(query.Limit))
	findOptions.SetSort(query.Sort)

	cursor, err := collection.Find(ctx, query.Filter, findOptions)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	docs, err := collection.CountDocuments(ctx, query.Filter)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	return cursor, *criteria.GeNextPage(docs), docs, nil

}

func (repository *Repository) Persist(data interface{}, filter *bson.D) error {

	if filter != nil {
		return repository.UpdateData(data, filter)
	}

	_, err := repository.GetCollection().InsertOne(context.TODO(), data)

	if err != nil {
		return err

	}

	return nil

}

func (repository *Repository) UpdateData(data interface{}, filter *bson.D) error {
	opts := options.Update().SetUpsert(true)
	updateData, err := repository.MongoConvert.PrepareUpsert(data)
	if err != nil {
		return err
	}

	update := bson.D{
		{
			"$set", updateData,
		},
	}

	_, err = repository.GetCollection().UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		return err

	}

	return nil
}
