package systemusecase

import systemdomain "digital-bank/internal/system/domain"

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

func (s *SearchAppClient) Run(clientID string) systemdomain.Result[*systemdomain.AppClient] {
	c, err := s.rep.GetClientByClientID(clientID)

	if err != nil {
		return systemdomain.NewResult[*systemdomain.AppClient](nil, &systemdomain.ErrorMessage{
			HttpCode: 404,
			Message:  err.Error(),
		})
	}

	return systemdomain.NewResult[*systemdomain.AppClient](c, nil)
}
