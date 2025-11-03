package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashThePassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatal("The error in hashing password is  : ", err)
		return "", errors.New("Error in hashing the password")
	}
	return string(hashedPassword), nil
}
func VerifyHashPassword(password string, hashedPassword string) bool {

	verify := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return verify == nil
}
