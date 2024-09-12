package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"ankit.com/event_booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID") // get the userID from JWT

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error converting id in integer",
			"error":   err,
		})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Event might not exist",
			"error":   err,
		})
		return
	}

	registrationID, err := models.CreateRegistration(userID, eventID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "There is problem while registration",
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":         "Registration sucessfull",
		"registration id": *registrationID,
	})

}

func cancelRegistraction(context *gin.Context) {
	registrationID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error converting id in integer",
			"error":   err,
		})
		return
	}
	retrivedUserID, err := models.GetRegistrationByID(registrationID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "There is no registration with provided registration ID",
			"error":   err,
		})
		return
	}
	userIDFromJWT := context.GetInt64("userID")
	fmt.Println(userIDFromJWT)
	fmt.Println(*retrivedUserID)
	if *retrivedUserID != userIDFromJWT {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "You can't cancel the registration because you're not authorized or you've not registed the event",
			"error":   err,
		})
		return
	}
	err = models.DeleteRegistration(registrationID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to cancel the registration",
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Registration cancelled successfully",
	})

}

func getRegistrations(context *gin.Context) {
	userID := context.GetInt64("userID")
	fmt.Println(userID)

	listOfRegistrations, err := models.ListRegistrations(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to find the list of registration for this user or Invalid userid",
			"error":   err,
		})
	}
	context.JSON(http.StatusOK, listOfRegistrations)
}
