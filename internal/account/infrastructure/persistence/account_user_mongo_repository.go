package accountpersistence

import (
	"context"
	accountdomain "digital-bank/internal/account/domain"
	mongo "digital-bank/pkg/mongodb"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

type (
	AccountUserMongoRepository struct {
		repo *mongo.Repository
	}
)

func NewAccountUserMongoRepository() *AccountUserMongoRepository {
	return &AccountUserMongoRepository{mongo.NewMongoRepository("accounts")}
}

func (r *AccountUserMongoRepository) Save(u *accountdomain.AccountUser) error {
	newUser := bson.D{
		{"email", u.Email},
		{"name", u.GetName()},
		{"password", u.GetPassword()},
		{"accountId", u.GetAccountID()},
		{"status", accountdomain.REGISTERED},
		{"createdAt", u.GetCreatedAt()},
		{"company", u.Company.GetIdentifier()},
		{"transactionFee", u.Company.GetCommissionsDefault()},
	}

	_, err := r.repo.GetCollection().InsertOne(context.TODO(), newUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountUserMongoRepository) EmailExists(email string) (bool, error) {
	var data map[string]interface{}
	err := r.repo.GetCollection().FindOne(context.TODO(), bson.D{{"email", email}}).Decode(data)
	if errors.Is(err, mongodb.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *AccountUserMongoRepository) FindByEmail(email string) (*accountdomain.AccountUser, error) {
	var data map[string]interface{}
	err := r.repo.GetCollection().FindOne(context.TODO(), bson.D{{"email", email}}).Decode(data)
	if errors.Is(err, mongodb.ErrNoDocuments) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	instance := accountdomain.NewAccountUser()
	return instance.AccountUserFromPrimitive(data), nil
}

func (r *AccountUserMongoRepository) UpdatePassword(u *accountdomain.AccountUser) error {
	_, err := r.repo.GetCollection().UpdateOne(context.TODO(), bson.D{{"email", u.Email}}, bson.D{{"$set", bson.D{{"password", u.GetPassword()}}}})
	if err != nil {
		return err
	}

	return nil
}
