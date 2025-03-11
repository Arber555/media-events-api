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
	media, err := services.GetMediaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, media)
}
