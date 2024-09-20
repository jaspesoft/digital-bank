package systemdomain

import (
	"digital-bank/internal"
	"encoding/base64"
	"time"
)

const (
	AppClientStatusActive   AppClientStatus = "ACTIVE"
	AppClientStatusDisabled AppClientStatus = "DISABLED"
)

type (
	AppClientStatus string

	AppClient struct {
		clientID              string          `json:"clientId"`
		name                  string          `json:"name"`
		companyID             string          `json:"companyId"`
		secret                string          `json:"secret"`
		technologyProviderFee TransactionFee  `json:"technologyProviderFee"`
		commissions           TransactionFee  `json:"commissions"`
		status                AppClientStatus `json:"status"`
		createdAt             time.Time       `json:"createdAt"`
	}

	AppClientIdentifier struct {
		ClientID string `json:"clientId"`
		Name     string `json:"name"`
	}

	AppClientRepository interface {
		GetClientByClientID(companyID string) (*AppClient, error)
		Upsert(client *AppClient) error
	}
)

func NewClient(clintID EntityID, name string, commissions TransactionFee, technologyProviderFee TransactionFee) *AppClient {

	client := &AppClient{
		clientID:              clintID.GetID(),
		name:                  name,
		status:                AppClientStatusActive,
		commissions:           commissions,
		technologyProviderFee: technologyProviderFee,
	}

	client.GenerateNewCredentialsAPI()

	return client
}

func (c *AppClient) IsActive() bool {
	return c.status == AppClientStatusActive
}

func (c *AppClient) GetClientID() string {
	return c.clientID
}

func (c *AppClient) GetStatus() AppClientStatus {
	return c.status

}

func (c *AppClient) GetTokenAPI() string {
	data := c.companyID + ":" + c.secret
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (c *AppClient) GenerateNewCredentialsAPI() {
	c.companyID = internal.GenerateRandomString(16)
	c.secret = base64.StdEncoding.EncodeToString([]byte(internal.GenerateRandomString(32)))
}

func (c *AppClient) GetIdentifier() AppClientIdentifier {
	return AppClientIdentifier{
		ClientID: c.clientID,
		Name:     c.name,
	}
}

func (c *AppClient) GetCommissionsDefault() TransactionFee {
	return c.commissions
}

func (c *AppClient) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"clientId":              c.clientID,
		"name":                  c.name,
		"companyId":             c.companyID,
		"secret":                c.secret,
		"status":                c.status,
		"commissions":           c.commissions.ToMap(),
		"technologyProviderFee": c.technologyProviderFee.ToMap(),
	}
}

func (c *AppClient) Disable() {
	c.status = AppClientStatusDisabled
}

func AppClientFromPrimitive(client map[string]interface{}) *AppClient {
	return &AppClient{
		clientID:              client["clientId"].(string),
		name:                  client["Name"].(string),
		companyID:             client["companyId"].(string),
		secret:                client["secret"].(string),
		status:                AppClientStatus(client["status"].(string)),
		commissions:           *TransactionFeeFromPrimitive(client["commissions"].(map[string]interface{})),
		technologyProviderFee: *TransactionFeeFromPrimitive(client["technologyProviderFee"].(map[string]interface{})),
		createdAt:             client["createdAt"].(time.Time),
	}
}
