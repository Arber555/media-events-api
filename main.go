package main

import (
	"media-events-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r)

	r.Run(":8080") // Start server on port 8080
}
