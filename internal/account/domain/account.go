package accountdomain

import (
	systemdomain "digital-bank/internal/system/domain"
	"digital-bank/internal/system/domain/criteria"
	"time"
)

const (
	REGISTERED        AccountStatus = "REGISTERED"
	CHANGES_REQUESTED AccountStatus = "CHANGES_REQUESTED"
	APPROVED          AccountStatus = "APPROVED"
	REJECTED          AccountStatus = "REJECTED"
	SUBMITTED         AccountStatus = "SUBMITTED"
	PROCESSING        AccountStatus = "PROCESSING"
	FROZEN            AccountStatus = "FROZEN"
	SUSPECTED_FRAUD   AccountStatus = "SUSPECTED_FRAUD"
)

type (
	AccountStatus string

	Account struct {
		AccountID      string                           `bson:"accountId" json:"accountId"`
		ApplicationID  string                           `bson:"applicationId" json:"applicationId"`
		Name           string                           `bson:"name" json:"name"`
		TypeAccount    AccountType                      `bson:"type" json:"clientType"`
		AccountHolder  AccountHolder                    `bson:"accountHolder" json:"accountHolder"`
		Status         AccountStatus                    `bson:"status" json:"status"`
		TransactionFee *systemdomain.TransactionFee     `bson:"transactionFee" json:"transactionFee"`
		ApprovedAt     time.Time                        `bson:"createdAt" json:"approvedAt"`
		OwnerRecord    systemdomain.AppClientIdentifier `bson:"clientOwnerRecord" json:"ownerRecord"`
	}

	AccountRepository interface {
		Paginate(criteria *criteria.Criteria) (criteria.Paginate, error)
		Upsert(account *Account) error
	}

	AccountProviderService interface {
		CreateApplication(a *Account) error
	}
)

func NewAccount(accountUser *AccountUser, accountHolder AccountHolder) *Account {
	err := accountHolder.SetAccountHolder(accountHolder)

	if err != nil {
		panic(err)
	}

	return &Account{
		AccountID:     accountUser.GetAccountID(),
		Name:          accountHolder.GetName(),
		TypeAccount:   accountHolder.GetType(),
		AccountHolder: accountHolder,
	}
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) GetAccountID() string {
	return a.AccountID
}

func (a *Account) GetType() AccountType {
	return a.AccountHolder.GetType()
}

func (a *Account) GetAccountHolder() AccountHolder {
	return a.AccountHolder
}

func (a *Account) GetStatus() AccountStatus {
	return a.Status
}

func (a *Account) GetTransactionFee() *systemdomain.TransactionFee {
	return a.TransactionFee
}

func (a *Account) ApproveAccount() {
	a.Status = APPROVED
	a.ApprovedAt = time.Now()
}

func (a *Account) RejectAccount() {
	a.Status = REJECTED
}

func (a *Account) RequestChanges() {
	a.Status = CHANGES_REQUESTED
}

func (a *Account) FreezeAccount() {
	a.Status = FROZEN
}

func (a *Account) SuspectFraud() {
	a.Status = SUSPECTED_FRAUD
}

func (a *Account) Processing() {
	a.Status = PROCESSING
}

func (a *Account) Submit() {
	a.Status = SUBMITTED
}

func (a *Account) SetApplicationID(applicationID string) {
	a.AccountID = applicationID
}

func (a *Account) GetApplicationID() string {
	return a.ApplicationID
}
