package users

import (
	"FinalProject/models"
	"FinalProject/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type (
	LoginRequest struct {
		Username string `json:"username,omitempty" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	LoginResponse struct {
		InsertedId uint `json:"insertId"`
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
		type Login struct {
			Username string `json:"username,omitempty"`
		}

		loginParams := Login{}
		ctx.ShouldBindJSON(&loginParams)
		user := models.User{}
		h.Db.Find(&user, loginParams.Username)

		// if loginParams.Username == "mike" || loginParams.Username == "rama" {
		if check := utils.CheckPasswordHash(user.Password, "abc"); check == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user": loginParams.Username,
				"nbf":  time.Date(2018, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
			})

			tokenStr, err := token.SignedString([]byte("supersaucysecret"))
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
			Message: loginParams,
		})
	}
}
