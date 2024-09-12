package routes

import (
	"fmt"
	"net/http"

	"ankit.com/event_booking/models"
	"ankit.com/event_booking/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
			"error":   err,
		})
		return
	}

	userid, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to save the data",
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"User ID": userid,
	})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameter",
			"error":   err,
		})
		return
	}
	isVerified := user.Verify()

	if !isVerified {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is WRONG",
		})
		return
	}
	fmt.Println(user.Email_ID, user.ID)
	token, err := utils.GenerateToken(user.Email_ID, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not able to authenticate, maybe problem with JWT",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Correct password",
		"token":   token,
	})
}
