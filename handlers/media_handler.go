package handlers

import (
	"media-events-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMedia(c *gin.Context) {
	query := c.Query("q")
	mediaList := services.SearchMedia(query)
	c.JSON(http.StatusOK, mediaList)
}

func GetMediaByID(c *gin.Context) {
	id := c.Param("id")
	media, found := services.GetMediaByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	c.JSON(http.StatusOK, media)
}
