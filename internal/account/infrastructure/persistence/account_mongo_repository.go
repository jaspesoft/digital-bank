package accountpersistence

import (
	"context"
	accountdomain "digital-bank/internal/account/domain"
	"digital-bank/internal/system/domain/criteria"
	mongo "digital-bank/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	AccountMongoRepository struct {
		repo *mongo.Repository
	}
)

func NewAccountMongoRepository() *AccountMongoRepository {
	return &AccountMongoRepository{
		repo: mongo.NewMongoRepository("account"),
	}
}

func (r *AccountMongoRepository) Paginate(cri *criteria.Criteria) (criteria.Paginate, error) {
	mongoPaginate, err := r.repo.SearchByCriteria(cri)

	if err != nil {
		return criteria.Paginate{}, err
	}

	var accounts []accountdomain.Account

	if err = mongoPaginate.Cursor.All(context.TODO(), &accounts); err != nil {
		return criteria.Paginate{}, err
	}

	return criteria.Paginate{
		Results:  accounts,
		NextPage: mongoPaginate.NextPage,
		Count:    mongoPaginate.Count,
		PrevPage: mongoPaginate.PrevPage,
	}, nil
}

func (r *AccountMongoRepository) Save(account accountdomain.Account) error {
	return r.repo.Persist(account, &bson.D{
		{"accountId", account.GetAccountID()},
	})
}
