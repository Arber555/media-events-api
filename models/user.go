package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
