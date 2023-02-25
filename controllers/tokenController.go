package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssss.tantalum/my-restful-api/auth"
	"github.com/ssss.tantalum/my-restful-api/database"
	"github.com/ssss.tantalum/my-restful-api/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
