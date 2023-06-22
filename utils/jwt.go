package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Create(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token will be expired after 12 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
