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
		Name           string                           `bson:"name" json:"name"`
		TypeAccount    AccountType                      `bson:"type" json:"clientType"`
		AccountHolder  AccountHolder                    `bson:"accountHolder" json:"accountHolder"`
		Status         AccountStatus                    `bson:"status" json:"status"`
		TransactionFee *systemdomain.TransactionFee     `bson:"transactionFee" json:"transactionFee"`
		CreatedAt      time.Time                        `bson:"createdAt" json:"createdAt"`
		ApprovedAt     time.Time                        `bson:"createdAt" json:"approvedAt"`
		OwnerRecord    systemdomain.AppClientIdentifier `bson:"clientOwnerRecord" json:"ownerRecord"`
	}

	AccountRepository interface {
		Paginate(criteria *criteria.Criteria) (criteria.Paginate, error)
	}

	AccountProviderService interface {
		CreateApplication(a Account) (string, error)
	}
)

func NewAccount(accountID systemdomain.EntityID, accountHolder AccountHolder, ownerRecord systemdomain.AppClient) *Account {
	accountHolder.SetAccountHolder(accountHolder)

	return &Account{
		AccountID:      accountID.GetID(),
		Name:           accountHolder.GetName(),
		TypeAccount:    accountHolder.GetType(),
		AccountHolder:  accountHolder,
		Status:         REGISTERED,
		TransactionFee: ownerRecord.GetCommissionsDefault(),
		OwnerRecord:    ownerRecord.GetIdentifier(),
		CreatedAt:      time.Now(),
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

func (a *Account) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":           a.Name,
		"type":           a.TypeAccount,
		"accountHolder":  a.AccountHolder.ToMap(),
		"status":         a.Status,
		"transactionFee": a.TransactionFee.ToMap(),
		"createdAt":      a.CreatedAt,
		"approvedAt":     a.ApprovedAt,
	}
}
