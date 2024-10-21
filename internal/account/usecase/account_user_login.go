package usecaseaccount

import (
	accountdomain "digital-bank/internal/account/domain"
	systemdomain "digital-bank/internal/system/domain"
	"log"
)

type (
	LoginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	AccountUserLogin struct {
		userRepository accountdomain.AccountUserRepository
		passAdapter    accountdomain.HashPasswordAdapter
		adapterToken   AdapterGenerateToken
	}

	AuthorizationResponse struct {
		AccessToken string `json:"accessToken"`
	}

	AdapterGenerateToken interface {
		CreateToken(user *accountdomain.TokenData) (string, error)
	}
)

func NewAccountUserLogin(
	repo accountdomain.AccountUserRepository, passAdapter accountdomain.HashPasswordAdapter,
	adapterToken AdapterGenerateToken,
) *AccountUserLogin {
	return &AccountUserLogin{
		userRepository: repo,
		passAdapter:    passAdapter,
		adapterToken:   adapterToken,
	}
}

func (u *AccountUserLogin) Run(req LoginReq) systemdomain.Result[*AuthorizationResponse] {
	user, err := u.userRepository.FindByEmail(req.Email)

	if err != nil {
		log.Println("find user by email in user account login error", err)
		return systemdomain.NewResult[*AuthorizationResponse](nil, systemdomain.NewError(400, "The email address or password is incorrect"))
	}

	if !u.passAdapter.ComparePassword(user.GetPassword(), req.Password) {
		return systemdomain.NewResult[*AuthorizationResponse](nil, systemdomain.NewError(400, "The password is incorrect"))
	}

	token, err := u.adapterToken.CreateToken(&accountdomain.TokenData{
		AccountID:      user.GetAccountID(),
		Status:         user.GetStatus(),
		TransactionFee: user.GetTransactionFee(),
	})

	if err != nil {
		log.Println("create token in user account login error", err)
		return systemdomain.NewResult[*AuthorizationResponse](nil, systemdomain.NewError(500, "internal server error in create token"))
	}

	return systemdomain.NewResult[*AuthorizationResponse](&AuthorizationResponse{AccessToken: token}, nil)
}
