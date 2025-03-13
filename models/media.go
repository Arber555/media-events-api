package models

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Transcript  string    `json:"transcript"`
	URL         string    `json:"url"`
	MediaType   string    `json:"mediaType"` // Image, Video, or Audio
	Topics      []string  `json:"topics"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
