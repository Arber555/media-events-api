package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID    `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Participants []Participant `json:"participants"`
	Date         time.Time       `json:"date"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type EventRegistration struct {
    UserID  uuid.UUID `json:"user_id"`  
    EventID uuid.UUID `json:"event_id"` 
    Role    string    `json:"role"`     
}

type Participant struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string `json:"role"`
}
