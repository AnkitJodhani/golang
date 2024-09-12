package middlewares

import (
	"net/http"

	"ankit.com/event_booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token needed!!",
		})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unable to verify your token",
			"error":   err,
		})
		return
	}
	context.Set("userID", *userID) // you can access userID value  by the key name "userID" & it is accessible whereever you have gin
	context.Next()
}
