package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhphan97/jwt-authentication-golang/database"
	"github.com/thinhphan97/jwt-authentication-golang/models"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)

	if record.Error != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
