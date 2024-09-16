package clientdomain

import (
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
		Name           string         `json:"name"`
		Type           AccountType    `json:"clientType"`
		AccountHolder  AccountHolder  `json:"accountHolder"`
		Status         AccountStatus  `json:"status"`
		TransactionFee TransactionFee `json:"transactionFee"`
		CreatedAt      time.Time      `bson:"createdAt" json:"createdAt"`
		ApprovedAt     time.Time      `bson:"createdAt" json:"approvedAt"`
	}
)

func NewAccount(accountHolder AccountHolder, transactionFee TransactionFee) *Account {
	accountHolder.SetAccountHolder(accountHolder)

	return &Account{
		Name:           accountHolder.GetName(),
		Type:           accountHolder.GetType(),
		AccountHolder:  accountHolder,
		Status:         REGISTERED,
		TransactionFee: transactionFee,
		CreatedAt:      time.Now(),
	}
}

func (a *Account) GetName() string {
	return a.Name
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

func (a *Account) GetTransactionFee() TransactionFee {
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

func (a *Account) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":           a.Name,
		"type":           a.Type,
		"accountHolder":  a.AccountHolder.ToMap(),
		"status":         a.Status,
		"transactionFee": a.TransactionFee.ToMap(),
		"createdAt":      a.CreatedAt,
		"approvedAt":     a.ApprovedAt,
	}
}
