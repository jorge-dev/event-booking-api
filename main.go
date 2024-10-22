package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/models"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":1234")
}

// Function to get the events
func getEvents(c *gin.Context) {
	events := models.GetAll()
	c.JSON(http.StatusOK, events)
}

// Function to create an event
func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data was provided"})
		return
	}

	event.ID = 1
	event.UserId = 1
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
