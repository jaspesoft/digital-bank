package systempersistence

import (
	"context"
	systemdomain "digital-bank/internal/system/domain"
	mongo "digital-bank/pkg/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	SystemMongoRepository struct {
		rep *mongo.Repository
	}
)

func NewSystemMongoRepository() *SystemMongoRepository {
	return &SystemMongoRepository{
		rep: mongo.NewMongoRepository("system_client"),
	}
}

func (r *SystemMongoRepository) Upsert(client *systemdomain.AppClient) error {
	return r.rep.Persist(client.ToMap(), &bson.D{
		{"accountId", client.GetClientID()},
	})
}

func (r *SystemMongoRepository) GetClientByClientID(companyID string) (*systemdomain.AppClient, error) {

	filter := bson.D{
		{"clientId", companyID},
	}

	var account map[string]interface{}

	err := r.rep.GetCollection().FindOne(context.TODO(), filter).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("Client not found")
		}
		return nil, err
	}

	return systemdomain.AppClientFromPrimitive(account), nil
}
