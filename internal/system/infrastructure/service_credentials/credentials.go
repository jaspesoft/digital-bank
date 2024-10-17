package credentials

import (
	systemdomain "digital-bank/internal/system/domain"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
)

type (
	Layer2Credentials struct {
		AuthToken string `json:"token"`
		URL       string `json:"url"`
		Signature string `json:"signatureKey"`
	}
	ServiceCredentials struct {
		CompanyID string `json:"companyId"`
		Layer2Credentials
	}
)

func FindApplicationClientCredentials(c *gin.Context) (*ServiceCredentials, error) {
	client, _ := c.Get("AppClient")
	appClient := client.(*systemdomain.AppClient)

	var clients []ServiceCredentials
	err := json.Unmarshal([]byte(os.Getenv("SERVICE_CREDENTIALS")), &clients)
	if err != nil {
		return nil, err
	}

	for _, client := range clients {
		if client.CompanyID == appClient.Credentials.CompanyID {
			return &client, nil
		}
	}

	return nil, errors.New("credentials not found")
}

func SearchApplicationClient(c *gin.Context) *systemdomain.AppClient {
	appClient, _ := c.Get("AppClient")
	return appClient.(*systemdomain.AppClient)

}
