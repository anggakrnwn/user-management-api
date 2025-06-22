package helpers

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
