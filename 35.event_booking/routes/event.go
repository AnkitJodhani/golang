package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"ankit.com/event_booking/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, _ := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error converting id in integer",
			"error":   err,
		})
		return
	}
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get the Event details",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Thare you are!!!",
		"error":   event,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	fmt.Println(event)
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't able to parse the data",
			"error":   err,
		})
		return
	}
	data, err := event.Save()
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"message": "Not able to save your event",
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Created successfully",
		"event":   data,
	})
}

func updateEvent(context *gin.Context) {
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
			"message": "Unable to get the Event details",
			"error":   err,
		})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't able to parse the data",
			"error":   err,
		})
		return
	}
	event.ID = eventID
	err = event.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error while updating events",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event has updated sucessfully",
	})
}

func deleteEvent(context *gin.Context) {
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

	var event models.Event
	event.ID = eventID
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error while deleting event",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event has deleted sucessfully",
	})

}
