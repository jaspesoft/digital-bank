package usecaseaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	systemdomain "digital-bank/internal/system/domain"
	"log"
)

type (
	AccountUserChangePassword struct {
		userRepository accountdomain.AccountUserRepository
		passAdapter    accountdomain.HashPasswordAdapter
	}

	ChangePasswordReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewAccountUserChangePassword(userRepository accountdomain.AccountUserRepository, passAdapter accountdomain.HashPasswordAdapter) *AccountUserChangePassword {
	return &AccountUserChangePassword{userRepository, passAdapter}
}

func (u *AccountUserChangePassword) Run(req ChangePasswordReq) systemdomain.Result[string] {
	user, err := u.userRepository.FindByEmail(req.Email)
	if err != nil {
		log.Println("find user by email in user account change password error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(404, "User not found"))
	}

	passEncrypt, _ := u.passAdapter.HashPassword(req.Password)
	user.SetPassword(passEncrypt)

	err = u.userRepository.UpdatePassword(user)
	if err != nil {
		log.Println("update password in user account change password error", err)
		return systemdomain.NewResult[string]("", systemdomain.NewError(500, "Internal server error in user account change password"))
	}

	return systemdomain.NewResult[string]("", nil)

}
