package systempersistence

import (
	"context"
	systemdomain "digital-bank/internal/system/domain"
	mongo "digital-bank/pkg/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	AppClientMongoRepository struct {
		rep *mongo.Repository
	}
)

func NewAppClientMongoRepository() *AppClientMongoRepository {
	return &AppClientMongoRepository{
		rep: mongo.NewMongoRepository("company_clients"),
	}
}

func (rc *AppClientMongoRepository) Upsert(appClient *systemdomain.AppClient) error {
	return rc.rep.Persist(*appClient, &bson.D{
		{"clientId", appClient.GetClientID()},
	})
}

func (rc *AppClientMongoRepository) GetClientByClientID(companyID string) (*systemdomain.AppClient, error) {

	filter := bson.D{
		{"clientId", companyID},
	}

	var account map[string]interface{}

	err := rc.rep.GetCollection().FindOne(context.TODO(), filter).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("Client not found")
		}
		return nil, err
	}

	return systemdomain.AppClientFromPrimitive(account), nil
}
