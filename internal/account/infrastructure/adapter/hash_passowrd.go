package accountadapter

import "golang.org/x/crypto/bcrypt"

type HashPasswordAdapter struct {
}

func NewHashPasswordAdapter() *HashPasswordAdapter {
	return &HashPasswordAdapter{}
}

func (h *HashPasswordAdapter) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *HashPasswordAdapter) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
