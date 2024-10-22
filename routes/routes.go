package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	v1 := server.Group("/v1/api")
	{
		v1.GET("/events", GetEvents)
		v1.GET("/events/:id", GetEvent)
		v1.POST("/events", middleware.ExtractEventAttributes(), CreateEvent)

	}
}
