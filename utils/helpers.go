package utils

import (
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
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
