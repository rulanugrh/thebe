package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Printf("Cant generate hash password: %v", err)
	}

	return string(bytes)
}

func CheckPassword(password string) error {

	comparePassword := HashPassword(password)
	err := bcrypt.CompareHashAndPassword([]byte(comparePassword), []byte(password))
	if err != nil {
		log.Printf("Cant compare password: %v", err)
	}

	return nil
}