package usecaseaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	reqaccount "digital-bank/internal/account/infrastructure/http/requests"
	systemdomain "digital-bank/internal/system/domain"
	"log"
)

type (
	ApplicationAccount struct {
		accountRepository accountdomain.AccountRepository
		serviceProvider   accountdomain.AccountProviderService
	}
)

func NewApplicationAccount(
	accountRepository accountdomain.AccountRepository,
	serviceProvider accountdomain.AccountProviderService,
) *ApplicationAccount {
	return &ApplicationAccount{
		accountRepository: accountRepository,
		serviceProvider:   serviceProvider,
	}
}

func (a *ApplicationAccount) Run(
	AccountUser *accountdomain.AccountUser, req reqaccount.ApplicationAccountCompanyRequest,
) systemdomain.Result[string] {

	account := accountdomain.NewAccount(AccountUser, &req.Company)

	err := a.serviceProvider.CreateApplication(account)

	if err != nil {
		log.Println("create application error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(400, err.Error()))
	}

	err = a.accountRepository.Upsert(account)

	if err != nil {
		log.Println("save account error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(500, "internal server error in save account"))
	}

	return systemdomain.NewResult[string](account.GetApplicationID(), nil)
}
