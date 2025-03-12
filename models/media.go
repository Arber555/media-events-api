package models

import "github.com/google/uuid"

type Media struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Transcript  string    `json:"transcript"`
	URL         string    `json:"url"`
	MediaType   string    `json:"mediaType"` // Image, Video, or Audio
	Topics      []string  `json:"topics"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
