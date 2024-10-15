package usecaseaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	systemdomain "digital-bank/internal/system/domain"
	"log"
)

type (
	AccountUserRegister struct {
		userRepository accountdomain.AccountUserRepository

		passAdapter     accountdomain.HashPasswordAdapter
		entityIDAdapter systemdomain.EntityIDAdapter
		evtBus          systemdomain.EventBus
	}

	AccountUserReq struct {
		Email       string                    `json:"email"`
		FirstName   string                    `json:"firstName"`
		MiddleName  string                    `json:"middleName"`
		LastName    string                    `json:"lastName"`
		AccountType accountdomain.AccountType `json:"userType"`
	}
)

func NewAccountUserRegister(
	repo accountdomain.AccountUserRepository,
	passAdapter accountdomain.HashPasswordAdapter,
	entityIDAdapter systemdomain.EntityIDAdapter,
	evtBus systemdomain.EventBus,
) *AccountUserRegister {
	return &AccountUserRegister{
		userRepository:  repo,
		passAdapter:     passAdapter,
		entityIDAdapter: entityIDAdapter,
		evtBus:          evtBus,
	}
}

func (u *AccountUserRegister) Run(req AccountUserReq, ownerRecord systemdomain.AppClient) systemdomain.Result[string] {

	exists, err := u.userRepository.EmailExists(req.Email)
	if err != nil {
		log.Println("validate exists email in user account register error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(500, "internal server error in user register"))
	}

	if exists {
		return systemdomain.NewResult[string]("", systemdomain.NewError(400, "The email address reported already exists on the platform"))
	}

	user := accountdomain.NewAccountUser()
	user.CreateOnboarding(req.Email, req.FirstName, req.MiddleName, req.LastName, ownerRecord)
	user.SetAccountID(u.entityIDAdapter)
	tmpPass := user.GeneratePassword(u.passAdapter)

	err = u.userRepository.Save(user)
	if err != nil {
		log.Println("save user account register error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(500, "Internal error when saving the information."))
	}

	_ = u.evtBus.Emit(map[string]string{
		"pass":   tmpPass,
		"client": user.GetName(),
	}, systemdomain.TOPIC_SENDMAIL)

	return systemdomain.NewResult[string]("An email will be sent with the access data to the platform.", nil)
}

func AccountUserFromPrimitive(d map[string]interface{}) {

}
