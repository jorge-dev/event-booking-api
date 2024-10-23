package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/models"
)

func RegisterForEvents(c *gin.Context) {
	// Get the event ID from the URL
	// Extract the user ID from the context
	// Call the models.RegisterForEvent function
	// Return the response

	userId := c.GetInt64("userId")
	if userId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User ID not found in context"})
		return
	}

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	eventFromDb, err := models.GetByID(eventId)
	if err != nil {
		errorMessage := "Could not find event with ID: " + strconv.FormatInt(eventId, 10)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errorMessage, "error": err.Error()})
		return
	}

	err = eventFromDb.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error registering for event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event", "event": eventFromDb})

}
func CancelRegistration(c *gin.Context) {

	userId := c.GetInt64("userId")
	if userId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User ID not found in context"})
		return
	}

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	eventFromDb, err := models.GetByID(eventId)
	if err != nil {
		errorMessage := "Could not find event with ID: " + strconv.FormatInt(eventId, 10)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errorMessage, "error": err.Error()})
		return
	}

	err = eventFromDb.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error cancelling registration for event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration for event", "event": eventFromDb})
}
