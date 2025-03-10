package services

import (
	"media-events-api/models"

	"github.com/google/uuid"
)

var users = []models.User{
	{ID: uuid.New(), Name: "John", LastName: "Doe", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: uuid.New(), Name: "Jane", LastName: "Doe", CreatedAt: "2025-01-02T00:00:00Z", UpdatedAt: "2025-01-02T00:00:00Z"},
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id string) (models.User, bool) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return models.User{}, false
	}

	for _, user := range users {
		if user.ID == parsedID {
			return user, true
		}
	}
	return models.User{}, false
}
