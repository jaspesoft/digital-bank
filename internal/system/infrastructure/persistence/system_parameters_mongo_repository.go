package systempersistence

import (
	"context"
	systemdomain "digital-bank/internal/system/domain"
	mongo "digital-bank/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	SystemParametersMongoRepository struct {
		rep *mongo.Repository
	}
)

func NewSystemParametersMongoRepository() *SystemParametersMongoRepository {
	return &SystemParametersMongoRepository{
		rep: mongo.NewMongoRepository("system_parameters"),
	}
}

func (sys *SystemParametersMongoRepository) GetSystemParameters() (*systemdomain.SystemParameters, error) {

	var systemParameters map[string]interface{}

	err := sys.rep.GetCollection().FindOne(context.TODO(), bson.D{}).Decode(&systemParameters)
	if err != nil {
		return nil, err
	}

	return systemdomain.SystemParametersFromPrimitive(systemParameters), nil
}
