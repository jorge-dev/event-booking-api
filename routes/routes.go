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
		v1.PUT("/events/:id", middleware.ExtractEventAttributes(), UpdateEvent)
		v1.DELETE("/events/:id", DeleteEvent)

		// User routes
		v1.POST("/signup", middleware.ExtractUserAttributes(), SignUp)
		v1.POST("/login", middleware.ExtractAuthUserAttributes(), Login)

	}
}
