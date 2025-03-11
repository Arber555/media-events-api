package handlers

import (
	"media-events-api/services"
	"net/http"
	"media-events-api/models"
	"github.com/google/uuid"
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

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	eventID, err := uuid.Parse(req.EventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	reg := models.EventRegistration{
		UserID:  userID, 
		EventID: eventID, 
		Role:    req.Role,
	}

	err = services.RegisterUserForEvent(reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
