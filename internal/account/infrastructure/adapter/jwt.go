package accountadapter

import (
	accountdomain "digital-bank/internal/account/domain"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type (
	ClaimsToken struct {
		AccountHolder accountdomain.TokenData `json:"user"`
		jwt.StandardClaims
	}
)

func CreateToken(data accountdomain.TokenData) (string, error) {
	expireTime := time.Now().Add(time.Hour * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":           expireTime,
		"accountHolder": data,
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
