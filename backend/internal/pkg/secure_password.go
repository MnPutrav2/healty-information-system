package pkg

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(pass string) (string, error) {
	pepper := os.Getenv("SECRET_STRING")
	secure := pass + pepper

	hash, err := bcrypt.GenerateFromPassword([]byte(secure), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
