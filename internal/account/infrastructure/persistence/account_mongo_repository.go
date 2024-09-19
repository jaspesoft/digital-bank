package persistence

import (
	"context"
	accountdomain "digital-bank/internal/account/domain"
	"digital-bank/internal/system/domain/criteria"
	mongo "digital-bank/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	AccountRepository struct {
		repo *mongo.Repository
	}
)

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		repo: mongo.NewMongoRepository("account"),
	}
}

func (r *AccountRepository) Paginate(cri *criteria.Criteria) (criteria.Paginate, error) {
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

func (r *AccountRepository) Save(account accountdomain.Account) error {
	return r.repo.Persist(account.ToMap(), &bson.D{
		{"accountId", account.GetAccountID()},
	})
}
