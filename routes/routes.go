package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	v1Public := server.Group("/v1/api")
	{
		v1Public.GET("/events", GetEvents)
		v1Public.GET("/events/:id", GetEvent)
		// User routes
		v1Public.POST("/signup", middleware.ExtractUserAttributes(), SignUp)
		v1Public.POST("/login", middleware.ExtractAuthUserAttributes(), Login)
	}

	v1Auth := server.Group("/v1/api")
	v1Auth.Use(middleware.Authenticate())
	{
		v1Auth.POST("/events", middleware.ExtractEventAttributes(), CreateEvent)
		v1Auth.PUT("/events/:id", middleware.ExtractEventAttributes(), UpdateEvent)
		v1Auth.DELETE("/events/:id", DeleteEvent)

		// registration routes
		v1Auth.POST("/events/:id/register", RegisterForEvents)
		v1Auth.DELETE("/events/:id/register", CancelRegistration)

	}

}
