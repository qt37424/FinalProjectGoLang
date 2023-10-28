package users

import (
	"FinalProject/models"
	"FinalProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
			if token, err := utils.CreateToken(&user); err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{"token": token})
				return
			}
		}

		ctx.JSON(http.StatusBadRequest, UnsignedResponse{
			Message: "wrong username or password",
		})
	}
}
