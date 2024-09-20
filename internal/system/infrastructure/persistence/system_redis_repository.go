package systempersistence

import (
	systemdomain "digital-bank/internal/system/domain"
	"digital-bank/pkg/cache"
	"errors"
)

type (
	SystemRedisRepository struct {
	}
)

func NewSystemRedisRepository() *SystemRedisRepository {
	return &SystemRedisRepository{}
}

func (r *SystemRedisRepository) GetClientByClientID(companyID string) (*systemdomain.AppClient, error) {
	data, err := cache.RecoverData(companyID)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Client not found")
	}

	return systemdomain.AppClientFromPrimitive(data), nil

}

func (r *SystemRedisRepository) Upsert(client *systemdomain.AppClient) error {
	panic("not implemented for redis") // TODO: Implement
	return nil
}
