package systempersistence

import (
	systemdomain "digital-bank/internal/system/domain"
	"digital-bank/pkg/cache"
	"errors"
)

type (
	AppClientRedisRepository struct {
	}
)

func NewAppClientRedisRepository() *AppClientRedisRepository {
	return &AppClientRedisRepository{}
}

func (r *AppClientRedisRepository) GetClientByClientID(companyID string) (*systemdomain.AppClient, error) {
	data, err := cache.RecoverData(companyID)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Client not found")
	}

	return systemdomain.AppClientFromPrimitive(data), nil

}

func (r *AppClientRedisRepository) Upsert(client *systemdomain.AppClient) error {
	panic("not implemented for redis") // TODO: Implement
	return nil
}

func (r *AppClientRedisRepository) GetClientByEmail(email string) (*systemdomain.AppClient, error) {
	panic("not implemented for redis") // TODO: Implement
	return nil, nil

}
