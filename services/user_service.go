package services

import (
	"errors"
	"media-events-api/models"

	"github.com/google/uuid"
)

var users = []models.User{
	{ID: uuid.MustParse("6b1d1f30-fcb2-11ec-b939-0242ac120002"), Name: "John", LastName: "Doe", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: uuid.MustParse("6b1d1f30-fcb2-11ec-b939-0242ac120003"), Name: "Jane", LastName: "Doe", CreatedAt: "2025-01-02T00:00:00Z", UpdatedAt: "2025-01-02T00:00:00Z"},
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id string) (*models.User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	for _, user := range users {
		if user.ID == parsedID {
			return &user, nil // Return a pointer instead of a struct
		}
	}
	return nil, errors.New("user not found")
}
