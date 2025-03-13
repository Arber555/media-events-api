package services

import (
	"errors"
	"strings"
	"time"
	"media-events-api/models"

	"github.com/google/uuid"
)

var mediaList = []models.Media{
	{
		ID:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		Title:       "AI in Healthcare",
		Description: "How AI is revolutionizing medicine.",
		Transcript:  "AI is used in diagnostics...",
		URL:         "http://example.com/ai-healthcare",
		MediaType:   "Video",
		Topics:      []string{"AI", "Health"},
		CreatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440001"),
		Title:       "Space Exploration",
		Description: "NASA's new space missions.",
		Transcript:  "NASA is launching a new rover...",
		URL:         "http://example.com/space",
		MediaType:   "Audio",
		Topics:      []string{"Space", "Science"},
		CreatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}

// Search and filter media based on query, topics, and mediaType
func SearchMedia(query, topic, mediaType string) []models.Media {
	var filtered []models.Media

	for _, media := range mediaList {
		// Check if the query matches the title or description
		matchesQuery := query == "" || strings.Contains(strings.ToLower(media.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(media.Description), strings.ToLower(query))

		// Check if the media matches the topic
		matchesTopic := topic == "" || containsIgnoreCase(media.Topics, topic)

		// Check if the mediaType matches
		matchesType := mediaType == "" || strings.EqualFold(media.MediaType, mediaType)

		// If all filters match, add to the results
		if matchesQuery && matchesTopic && matchesType {
			filtered = append(filtered, media)
		}
	}

	return filtered
}

// Helper function to check if a slice contains a value (case-insensitive)
func containsIgnoreCase(slice []string, value string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, value) {
			return true
		}
	}
	return false
}

func GetMediaByID(id string) (*models.Media, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	for _, media := range mediaList {
		if media.ID == parsedID {
			return &media, nil
		}
	}
	return nil, errors.New("media not found")
}
