package routes

import (
	"media-events-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine) {
	event := r.Group("/events")
	{
		event.GET("", handlers.GetEvents)
		event.GET("/:id", handlers.GetEventByID)
		event.POST("/register", handlers.RegisterForEvent)
	}
}
