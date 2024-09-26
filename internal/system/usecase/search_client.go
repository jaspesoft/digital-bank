package systemusecase

import (
	systemdomain "digital-bank/internal/system/domain"
)

type (
	SearchAppClient struct {
		rep systemdomain.AppClientRepository
	}
)

func NewSearchAppClient(repository systemdomain.AppClientRepository) *SearchAppClient {

	return &SearchAppClient{
		rep: repository,
	}
}

func (s *SearchAppClient) Run(companyID string) systemdomain.Result[*systemdomain.AppClient] {
	c, err := s.rep.GetClientByCompanyID(companyID)
	if err != nil {
		return systemdomain.NewResult[*systemdomain.AppClient](nil, systemdomain.NewError(
			404, err.Error(),
		))
	}

	return systemdomain.NewResult(c, nil)
}
