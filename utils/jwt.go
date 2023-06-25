package utils

import (
	"FinalProject/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("[ERROR] Signing string error: %v", err)
	}
	return tokenStr, nil
}

func ParsingToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenStr := ExtractToken(ctx)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[ERROR] Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	// bearerToken, err := ctx.Cookie("Authorization")
	// if err != nil || bearerToken == "" {
	// 	return ""
	// }
	return bearerToken
}
