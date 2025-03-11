package handlers

import (
	"media-events-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	location := c.Query("location")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	events := services.SearchEvents(location, dateFrom, dateTo)
	c.JSON(http.StatusOK, events)
}

func GetEventByID(c *gin.Context) {
	id := c.Param("id")
	event, found := services.GetEventByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func RegisterForEvent(c *gin.Context) {
	var req struct {
		UserID  string `json:"user_id"`
		EventID string `json:"event_id"`
		Role    string `json:"role"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.RegisterUserForEvent(req.UserID, req.EventID, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
