package routes

import (
	"media-events-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", handlers.GetUsers)
		users.GET("/:id", handlers.GetUserByID)
	}
}
