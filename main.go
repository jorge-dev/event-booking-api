package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jorge-dev/ev-book/db"
	"github.com/jorge-dev/ev-book/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	// Register the routes
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
