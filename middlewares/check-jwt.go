package middlewares

import (
	"FinalProject/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func CheckjwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.ParsingToken(ctx)
		if err != nil {
			fmt.Printf("[ERROR] JwtAuth error: %+v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": token})
			ctx.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			jsonCurrentUser, err := json.Marshal(claims["userId"])
			if err != nil {
				fmt.Printf("[ERROR] Get current user error: %+v", err)
			}
			var currentUser uint
			if err := json.Unmarshal(jsonCurrentUser, &currentUser); err != nil {
				fmt.Printf("[ERROR] Get current user error: %+v", err)
			}
			// ctx.JSON(http.StatusUnauthorized, gin.H{"error": currentUser})
			ctx.Set("currentUserId", currentUser)
		}
		ctx.Next()
	}
}
