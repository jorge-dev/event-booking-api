package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/models"
)

// Function to get the events
// GetEvents handles the HTTP request to retrieve all events.
// It fetches all events from the database using the models.GetAll function.
// If an error occurs during the retrieval, it responds with an HTTP 500 status code and an error message.
// On success, it responds with an HTTP 200 status code and the list of events.
func GetEvents(c *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Function to get an event
// GetEvent handles the HTTP request to retrieve an event by its ID.
// It expects an "id" parameter in the URL, which should be a valid integer.
// If the "id" parameter is invalid, it responds with a 400 Bad Request status and an error message.
// If the event retrieval fails due to a server error, it responds with a 500 Internal Server Error status and the error message.
// On success, it responds with a 200 OK status and the event data in JSON format.
func GetEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	event, err := models.GetByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Function to create an event
// CreateEvent handles the creation of a new event.
// It retrieves the event from the context, saves it to the database, and returns a JSON response.
// If the event is not found in the context or if there is an error during saving, it returns an error response.
//
// @param c *gin.Context - The Gin context which contains the request and response objects.
//
// @response 201 - Event created successfully with the event details.
// @response 500 - Internal server error with an error message.
func CreateEvent(c *gin.Context) {
	event, exists := c.Get("event")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found in context"})
		return
	}

	eventModel := event.(models.Event)
	eventModel.ID = 1
	eventModel.UserId = 1
	err := eventModel.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": eventModel})

}
