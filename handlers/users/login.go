package users

import (
	"FinalProject/models"
	"FinalProject/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type (
	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	LoginResponse struct {
		IdToken uint `json:"insertId"`
	}
	Login struct {
		Username string `uri:"username,omitempty"`
	}
	UnsignedResponse struct {
		Message interface{} `json:"message"`
	}
	SignedResponse struct {
		Token   string `json:"token"`
		Message string `json:"message"`
	}
)

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		loginParams := LoginRequest{}
		ctx.ShouldBindJSON(&loginParams)
		user := models.User{}
		h.Db.Where("username = ?", loginParams.Username).Find(&user)

		if check := utils.CheckPasswordHash(user.Password, loginParams.Password); check == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":     user.ID,
				"authorized": true,
				"nbf":        time.Date(2018, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
			})

			// create a complete, signed JWT
			// randStr := utils.RandomString(10)
			tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, UnsignedResponse{
					Message: err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, SignedResponse{
				Token:   tokenStr,
				Message: "logged in",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, UnsignedResponse{
			Message: "wrong username or password",
		})
	}
}
