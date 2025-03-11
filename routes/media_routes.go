package routes

import (
	"media-events-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterMediaRoutes(r *gin.Engine) {
	media := r.Group("/media")
	{
		media.GET("/", handlers.GetMedia)
		media.GET("/:id", handlers.GetMediaByID)
	}
}
