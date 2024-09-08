package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"ankit.com/temp/db"
	"ankit.com/temp/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("learning golang")
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "🚫 Invalid parameter",
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "📃 Listing out events",
		"events":  events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	event.UserID = rand.Intn(10000)
	err := context.ShouldBindJSON(&event)
	fmt.Println(event) // for testing
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "🚫 Invalid parameter",
			"error":   err,
		})
		return
	}
	result, err := event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "😟 Unable to save data in db",
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"message": "🔥 event has been created successfully",
		"event":   result,
	})
}
