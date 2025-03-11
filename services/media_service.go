package services

import (
	"media-events-api/models"
	"strings"

	"github.com/google/uuid"
)

var mediaList = []models.Media{
	{ID: uuid.New(), Title: "Media 1", Description: "Description 1", Transcript: "Transcript 1", URL: "http://example.com/1", CreatedAt: "2025-01-01T00:00:00Z", UpdatedAt: "2025-01-01T00:00:00Z"},
	{ID: uuid.New(), Title: "Media 2", Description: "Description 2", Transcript: "Transcript 2", URL: "http://example.com/2", CreatedAt: "2025-01-02T00:00:00Z", UpdatedAt: "2025-01-02T00:00:00Z"},
}

func SearchMedia(query string) []models.Media {
	if query == "" {
		return mediaList
	}
	filtered := []models.Media{}
	for _, media := range mediaList {
		if strings.Contains(strings.ToLower(media.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(media.Description), strings.ToLower(query)) {
			filtered = append(filtered, media)
		}
	}
	return filtered
}

func GetMediaByID(id string) (models.Media, bool) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return models.Media{}, false
	}

	for _, media := range mediaList {
		if media.ID == parsedID {
			return media, true
		}
	}
	return models.Media{}, false
}
