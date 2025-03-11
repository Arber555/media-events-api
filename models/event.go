package models

import "github.com/google/uuid"

type Event struct {
	ID           uuid.UUID    `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Participants []Participant `json:"participants"`
	Date         string       `json:"date"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    string       `json:"updated_at"`
}

type Participant struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
