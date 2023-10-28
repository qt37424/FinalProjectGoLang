package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ExtractUsernameFromToken(tokenString string) (uint, error) {
	var userId uint
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["nbf"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId64, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userId"]), 10, 64)
		if err == nil {
			return 0, fmt.Errorf("unexpected signing method: %v", token.Header["nbf"])
		}
		userId = uint(userId64)
	}

	return userId, nil
}
