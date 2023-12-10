package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Printf("Cannot hashing password, because: %s", err.Error())
	}

	return string(bytes)
}

func CheckPassword(password string, compare []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), compare)
	if err != nil {
		log.Printf("Password not same")
		return err
	}

	return nil
}