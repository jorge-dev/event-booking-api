package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/models"
)

func ExtractEventAttributes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Data struct {
				Attributes models.Event `json:"attributes"`
			} `json:"data"`
		}

		// Read the body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
			c.Abort()
			return
		}

		// Reset the request body so it can be read again
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Bind the JSON to the input struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data was provided", "error": err.Error()})
			c.Abort()
			return
		}

		input.Data.Attributes.CreatedAt = time.Now()
		// Set the event in the context
		c.Set("event", input.Data.Attributes)
		c.Next()
	}
}

func ExtractUserAttributes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Data struct {
				Attributes models.User `json:"attributes"`
			} `json:"data"`
		}

		// Read
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
			c.Abort()
			return
		}

		// Reset the request body so it can be read again
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Bind the JSON to the input struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data was provided", "error": err.Error()})
			c.Abort()
			return
		}

		// Set the user in the context
		input.Data.Attributes.CreatedAt = time.Now()
		c.Set("user", input.Data.Attributes)
		c.Next()
	}
}

func ExtractAuthUserAttributes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Data struct {
				Attributes models.AuthUser `json:"attributes"`
			} `json:"data"`
		}

		// Read
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body", "error": err.Error()})
			c.Abort()
			return
		}

		// Reset the request body so it can be read again
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Bind the JSON to the input struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Data was provided", "error": err.Error()})
			c.Abort()
			return
		}

		// Set the user in the context
		c.Set("user", input.Data.Attributes)
		c.Next()
	}
}
