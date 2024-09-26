package pkg

import "golang.org/x/crypto/bcrypt"

type (
	BcryptPassword struct {
	}
)

func NewBcryptPassword() *BcryptPassword {
	return &BcryptPassword{}
}

func (b *BcryptPassword) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (b *BcryptPassword) CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}
