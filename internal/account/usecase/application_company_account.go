package usecaseaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	reqaccount "digital-bank/internal/account/infrastructure/http/requests"
	systemdomain "digital-bank/internal/system/domain"
)

type (
	ApplicationAccount struct {
		accountRepository accountdomain.AccountRepository
	}
)

func NewApplicationAccount(accountRepository accountdomain.AccountRepository) *ApplicationAccount {
	return &ApplicationAccount{
		accountRepository: accountRepository,
	}
}

func (a *ApplicationAccount) Run(req reqaccount.ApplicationAccountCompanyRequest) systemdomain.Result[string] {

	return systemdomain.NewResult[string]("applicationId", nil)
}
