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
		accountID      string                           `bson:"accountId" json:"accountId"`
		name           string                           `bson:"name" json:"name"`
		typeAccount    AccountType                      `bson:"type" json:"clientType"`
		accountHolder  AccountHolder                    `bson:"accountHolder" json:"accountHolder"`
		status         AccountStatus                    `bson:"status" json:"status"`
		transactionFee systemdomain.TransactionFee      `bson:"transactionFee" json:"transactionFee"`
		createdAt      time.Time                        `bson:"createdAt" json:"createdAt"`
		approvedAt     time.Time                        `bson:"createdAt" json:"approvedAt"`
		ownerRecord    systemdomain.AppClientIdentifier `bson:"clientOwnerRecord" json:"ownerRecord"`
	}

	AccountRepository interface {
		Paginate(criteria *criteria.Criteria) (criteria.Paginate, error)
	}
)

func NewAccount(accountID systemdomain.EntityID, accountHolder AccountHolder, ownerRecord systemdomain.AppClient) *Account {
	accountHolder.SetAccountHolder(accountHolder)

	return &Account{
		accountID:      accountID.GetID(),
		name:           accountHolder.GetName(),
		typeAccount:    accountHolder.GetType(),
		accountHolder:  accountHolder,
		status:         REGISTERED,
		transactionFee: ownerRecord.GetCommissionsDefault(),
		ownerRecord:    ownerRecord.GetIdentifier(),
		createdAt:      time.Now(),
	}
}

func (a *Account) GetName() string {
	return a.name
}

func (a *Account) GetAccountID() string {
	return a.accountID
}

func (a *Account) GetType() AccountType {
	return a.accountHolder.GetType()
}

func (a *Account) GetAccountHolder() AccountHolder {
	return a.accountHolder
}

func (a *Account) GetStatus() AccountStatus {
	return a.status
}

func (a *Account) GetTransactionFee() systemdomain.TransactionFee {
	return a.transactionFee
}

func (a *Account) ApproveAccount() {
	a.status = APPROVED
	a.approvedAt = time.Now()
}

func (a *Account) RejectAccount() {
	a.status = REJECTED
}

func (a *Account) RequestChanges() {
	a.status = CHANGES_REQUESTED
}

func (a *Account) FreezeAccount() {
	a.status = FROZEN
}

func (a *Account) SuspectFraud() {
	a.status = SUSPECTED_FRAUD
}

func (a *Account) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":           a.name,
		"type":           a.typeAccount,
		"accountHolder":  a.accountHolder.ToMap(),
		"status":         a.status,
		"transactionFee": a.transactionFee.ToMap(),
		"createdAt":      a.createdAt,
		"approvedAt":     a.approvedAt,
	}
}
