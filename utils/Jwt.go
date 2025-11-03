package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}

var secretKey = []byte("hey-there-this-is-the-jwt")

func GenerateJwt(name string, email string) (string, error) {
	expiryTime := time.Now().Add(time.Hour * 24 * 10)
	claims := &Claims{
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)

}
