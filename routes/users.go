package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/models"
)

func SignUp(context *gin.Context) {
	userEvent, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User not found in context"})
		return
	}
	userModel := userEvent.(models.User)
	err := userModel.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// remove the password from the response
	userModel.Password = ""
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": userModel})

}

func Login(context *gin.Context) {
	userEvent, exists := context.Get("user")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User not found in context"})
		return
	}
	userModel := userEvent.(models.AuthUser)
	token, err := userModel.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
