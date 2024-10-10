package layer2

import (
	accountdomain "digital-bank/internal/account/domain"
	layer2helpers "digital-bank/pkg/services/layer2/helpers"
	"encoding/json"
	"log"
)

type (
	Layer2Application struct {
		c *Layer2
	}

	applicationAPIResponse struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
)

func NewLayer2Application() *Layer2Application {
	return &Layer2Application{
		c: NewLayer2(),
	}
}

func (l *Layer2Application) CreateApplication(a accountdomain.Account) (string, error) {
	l.c.payload = layer2helpers.ApplicationPayloadPrepare(a)
	l.c.endPointURL = "v1/applications"

	bodyResp, err := l.c.Post()
	if err != nil {
		log.Println("Error on create application", err)
		return "", err
	}

	var res applicationAPIResponse
	if err := json.Unmarshal(bodyResp, &res); err != nil {
		return "", err
	}

	return res.Data.ID, nil
}

func (l *Layer2Application) sendPartner(a accountdomain.Account, applicationId string) error {
	holder := a.GetAccountHolder().(accountdomain.CompanyAccountHolder)

	partners := holder.GetPartners()

	for _, partner := range partners {
		l.c.endPointURL = "v1/applications/" + applicationId + "/individual"
		l.c.payload = layer2helpers.PartnerPayloadPrepare(partner)

	}

	return nil
}
