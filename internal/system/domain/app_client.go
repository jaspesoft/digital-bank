package systemdomain

import (
	"digital-bank/internal"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	AppClientStatusActive   AppClientStatus = "ACTIVE"
	AppClientStatusDisabled AppClientStatus = "DISABLED"
)

type (
	AppClientStatus string

	AppClient struct {
		ClientID              string          `json:"clientId"`
		CompanyName           string          `json:"companyName"`
		Email                 string          `json:"email"`
		PhoneNumber           string          `json:"phoneNumber"`
		CompanyID             string          `json:"companyId"`
		Secret                string          `json:"secret"`
		TechnologyProviderFee *TransactionFee `json:"technologyProviderFee"`
		Commissions           *TransactionFee `json:"commissions"`
		Status                AppClientStatus `json:"status"`
		CreatedAt             time.Time       `json:"createdAt"`
	}

	AppClientIdentifier struct {
		ClientID    string `json:"clientId"`
		CompanyName string `json:"companyName"`
	}

	AppClientRepository interface {
		GetClientByCompanyID(companyID string) (*AppClient, error)
		GetClientByEmail(email string) (*AppClient, error)
		Upsert(client *AppClient) error
	}
)

func NewAppClient(
	clintID EntityID, companyName, email, phoneNumber string, commissions *TransactionFee, technologyProviderFee *TransactionFee,
) *AppClient {

	client := &AppClient{
		ClientID:              clintID.GetID(),
		CompanyName:           companyName,
		Email:                 email,
		PhoneNumber:           phoneNumber,
		Status:                AppClientStatusActive,
		Commissions:           commissions,
		TechnologyProviderFee: technologyProviderFee,
		CreatedAt:             time.Now().UTC(),
	}

	client.GenerateNewCredentialsAPI()

	return client
}

func (c *AppClient) IsActive() bool {
	return c.Status == AppClientStatusActive
}

func (c *AppClient) GetClientID() string {
	return c.ClientID
}

func (c *AppClient) GetStatus() AppClientStatus {
	return c.Status

}

func (c *AppClient) GetTokenAPI() string {
	data := c.CompanyID + ":" + c.Secret
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (c *AppClient) GenerateNewCredentialsAPI() {
	c.CompanyID = internal.GenerateRandomString(16)
	c.Secret = base64.StdEncoding.EncodeToString([]byte(internal.GenerateRandomString(32)))
}

func (c *AppClient) GetIdentifier() AppClientIdentifier {
	return AppClientIdentifier{
		ClientID:    c.ClientID,
		CompanyName: c.CompanyName,
	}
}

func (c *AppClient) GetCommissionsDefault() *TransactionFee {
	return c.Commissions
}

func (c *AppClient) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"clientId":              c.ClientID,
		"companyName":           c.CompanyName,
		"companyId":             c.CompanyID,
		"createdAt":             c.CreatedAt,
		"email":                 c.Email,
		"phoneNumber":           c.PhoneNumber,
		"secret":                c.Secret,
		"status":                c.Status,
		"commissions":           c.Commissions.ToMap(),
		"technologyProviderFee": c.TechnologyProviderFee.ToMap(),
	}
}

func (c *AppClient) Disable() {
	c.Status = AppClientStatusDisabled
}

func AppClientFromPrimitive(client map[string]interface{}) *AppClient {

	var myRawDate time.Time
	if rawDate, ok := client["createdAt"].(primitive.DateTime); ok {
		myRawDate = rawDate.Time() // Convertir a time.Time
	}

	return &AppClient{
		ClientID:              client["clientId"].(string),
		CompanyName:           client["companyName"].(string),
		CompanyID:             client["companyId"].(string),
		Secret:                client["secret"].(string),
		Status:                AppClientStatus(client["status"].(string)),
		Commissions:           TransactionFeeFromPrimitive(client["commissions"].(map[string]interface{})),
		TechnologyProviderFee: TransactionFeeFromPrimitive(client["technologyProviderFee"].(map[string]interface{})),
		CreatedAt:             myRawDate,
		Email:                 client["email"].(string),
		PhoneNumber:           client["phoneNumber"].(string),
	}
}
