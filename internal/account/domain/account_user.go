package accountdomain

import (
	"digital-bank/internal"
	systemdomain "digital-bank/internal/system/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	AccountUser struct {
		Email          string      `json:"email"`
		FirstName      string      `json:"firstName"`
		MiddleName     string      `json:"middleName"`
		LastName       string      `json:"lastName"`
		AccountType    AccountType `json:"userType"`
		Company        systemdomain.AppClient
		password       string                       `json:"password"`
		accountID      string                       ` json:"accountId"`
		createdAt      time.Time                    `json:"createdAt"`
		status         AccountStatus                `json:"status"`
		transactionFee *systemdomain.TransactionFee `json:"transactionFee"`
	}

	TokenData struct {
		AccountID      string                       `json:"accountId"`
		Status         AccountStatus                `json:"status"`
		TransactionFee *systemdomain.TransactionFee `json:"transactionFee"`
	}

	HashPasswordAdapter interface {
		HashPassword(password string) (string, error)
		ComparePassword(hashedPassword, password string) bool
	}

	AccountUserRepository interface {
		Save(u *AccountUser) error
		EmailExists(email string) (bool, error)
		FindByEmail(email string) (*AccountUser, error)
		UpdatePassword(u *AccountUser) error
	}
)

func NewAccountUser() *AccountUser {
	return &AccountUser{}
}

func (o *AccountUser) CreateOnboarding(email, firstName, middleName, lastName string, accountType AccountType, ownerRecord systemdomain.AppClient) {
	o.Email = email
	o.FirstName = firstName
	o.MiddleName = middleName
	o.LastName = lastName
	o.createdAt = time.Now()
	o.Company = ownerRecord
	o.AccountType = accountType
}

func (o *AccountUser) GetStatus() AccountStatus {
	return o.status
}

func (o *AccountUser) GetTransactionFee() *systemdomain.TransactionFee {
	return o.transactionFee
}

func (o *AccountUser) GeneratePassword(passAdapter HashPasswordAdapter) string {
	pass := internal.GenerateRandomString(8)
	o.password, _ = passAdapter.HashPassword(pass)

	return pass
}

func (o *AccountUser) SetAccountID(entityIDAdapter systemdomain.EntityIDAdapter) {
	o.accountID = entityIDAdapter.GetID()
}

func (o *AccountUser) GetAccountID() string {
	return o.accountID
}

func (o *AccountUser) GetCreatedAt() time.Time {
	return o.createdAt
}

func (o *AccountUser) GetPassword() string {
	return o.password
}

func (o *AccountUser) SetPassword(password string) {
	o.password = password
}

func (o *AccountUser) GetName() string {
	if o.AccountType == COMPANY_CLIENT {
		return o.FirstName
	}

	if o.MiddleName != "" {
		return o.FirstName + " " + o.MiddleName + " " + o.LastName
	}
	return o.FirstName + " " + o.LastName
}

func (o *AccountUser) AccountUserFromPrimitive(d map[string]interface{}) *AccountUser {

	var name, middleName, lastName string
	if d["type"].(string) == "COMPANY" {
		name = d["name"].(string)
	} else {
		name = d["firstName"].(string)
		middleName = d["middleName"].(string)
		lastName = d["lastName"].(string)
	}

	return &AccountUser{
		Email:          d["email"].(string),
		FirstName:      name,
		MiddleName:     middleName,
		LastName:       lastName,
		AccountType:    AccountType(d["type"].(string)),
		password:       d["password"].(string),
		accountID:      d["accountId"].(string),
		createdAt:      d["createdAt"].(primitive.DateTime).Time(),
		status:         AccountStatus(d["status"].(string)),
		transactionFee: convertToTransactionFee(d["transactionFee"].(map[string]interface{})),
	}
}

func convertToTransactionFee(data map[string]interface{}) *systemdomain.TransactionFee {

	domesticUSA := systemdomain.DomesticUSA{
		ACH: struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		}{
			IN:  data["domesticUsa"].(map[string]interface{})["ach"].(map[string]interface{})["in"].(float64),
			OUT: data["domesticUsa"].(map[string]interface{})["ach"].(map[string]interface{})["out"].(float64),
		},
		FedWire: struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		}{
			IN:  data["domesticUsa"].(map[string]interface{})["fedWire"].(map[string]interface{})["in"].(float64),
			OUT: data["domesticUsa"].(map[string]interface{})["fedWire"].(map[string]interface{})["out"].(float64),
		},
	}

	swift := systemdomain.SwiftUSA{
		IN:  data["swiftUsa"].(map[string]interface{})["in"].(float64),
		OUT: data["swiftUsa"].(map[string]interface{})["out"].(float64),
	}

	swap := systemdomain.Swap{
		Buy:  data["swap"].(map[string]interface{})["buy"].(float64),
		Sell: data["swap"].(map[string]interface{})["sell"].(float64),
	}
	return systemdomain.NewTransactionFee(
		domesticUSA,
		swift,
		swap,
	)
}
