package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashToken(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("cannot generate hash token: %v", err)
	}
	return string(bytes), nil
}
func CheckTokenHash(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err == nil
}
