package services

import (
	"errors"
	"media-events-api/models"
	"time"
	"github.com/google/uuid"
)

var users = []models.User{
	{ID: uuid.MustParse("6b1d1f30-fcb2-11ec-b939-0242ac120002"), Name: "John", LastName: "Doe", CreatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
	{ID: uuid.MustParse("6b1d1f30-fcb2-11ec-b939-0242ac120003"), Name: "Jane", LastName: "Doe", CreatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
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
