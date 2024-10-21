package accountdomain

import (
	"digital-bank/internal"
	systemdomain "digital-bank/internal/system/domain"
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

func (o *AccountUser) CreateOnboarding(email, firstName, middleName, lastName string, ownerRecord systemdomain.AppClient) {
	o.Email = email
	o.FirstName = firstName
	o.MiddleName = middleName
	o.LastName = lastName
	o.createdAt = time.Now()
	o.Company = ownerRecord
}

func (o *AccountUser) GetStatus() AccountStatus {
	return o.status
}

func (o *AccountUser) GetTransactionFee() *systemdomain.TransactionFee {
	return o.transactionFee
}

func (o *AccountUser) GeneratePassword(passAdapter HashPasswordAdapter) string {
	pass := internal.GenerateRandomString(24)
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
	return &AccountUser{
		Email:          d["email"].(string),
		FirstName:      d["firstName"].(string),
		MiddleName:     d["middleName"].(string),
		LastName:       d["lastName"].(string),
		password:       d["password"].(string),
		accountID:      d["accountId"].(string),
		createdAt:      d["createdAt"].(time.Time),
		status:         d["status"].(AccountStatus),
		transactionFee: d["transactionFee"].(*systemdomain.TransactionFee),
	}
}
