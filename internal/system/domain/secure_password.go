package systemdomain

type (
	SecurePassword interface {
		Hash(password string) (string, error)
		CheckPasswordHash(password, hashedPassword string) bool
	}
)
