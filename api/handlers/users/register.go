package users

import (
	"FinalProject/models"
	"FinalProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	RegisterRequest struct {
		Username string `json:"username" binding:"required"`
		Role     string `json:"role" binding:"required"` // if necessary convert role from string to enum
		Password string `json:"password" binding:"required"`
	}
	RegisterResponse struct {
		InsertedId uuid.UUID `json:"insertId"`
	}
)

func (h *UserHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &RegisterRequest{}
		if err := ctx.ShouldBindJSON(body); err != nil { // here we also have a method call MustBindJSON
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		password := utils.Hash(body.Password)
		user := &models.User{
			Username: body.Username,
			Password: password,
		}
		result := h.Db.Create(user)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &RegisterResponse{InsertedId: user.ID})
	}
}
