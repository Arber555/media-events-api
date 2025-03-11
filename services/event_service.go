package services

import (
	"errors"
	"media-events-api/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

// In-memory event storage
var eventList = []models.Event{
	{
		ID:          uuid.New(),
		Title:       "Spring Conference 2025",
		Description: "Join us for a Spring Conference of fellowship, worship, and spiritual renewal.",
		Date:        "2025-04-15T10:00:00Z",
		CreatedAt:   "2025-03-01T00:00:00Z",
		UpdatedAt:   "2025-03-01T00:00:00Z",
	},
	{
		ID:          uuid.New(),
		Title:       "Music Festival",
		Description: "Celebrate faith and music at our church uplifting Spring Music Festival!",
		Date:        "2025-07-10T15:00:00Z",
		CreatedAt:   "2025-03-02T00:00:00Z",
		UpdatedAt:   "2025-03-02T00:00:00Z",
	},
}

func SearchEvents(location, dateFrom, dateTo string) []models.Event {
	filtered := []models.Event{}

	for _, event := range eventList {
		if location != "" && !strings.Contains(strings.ToLower(event.Description), strings.ToLower(location)) {
			continue
		}

		eventDate, err := time.Parse(time.RFC3339, event.Date)
		if err != nil {
			continue
		}

		if dateFrom != "" {
			fromDate, _ := time.Parse(time.RFC3339, dateFrom)
			if eventDate.Before(fromDate) {
				continue
			}
		}

		if dateTo != "" {
			toDate, _ := time.Parse(time.RFC3339, dateTo)
			if eventDate.After(toDate) {
				continue
			}
		}

		filtered = append(filtered, event)
	}

	return filtered
}

func GetEventByID(id string) (models.Event, bool) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return models.Event{}, false
	}

	for _, event := range eventList {
		if event.ID == parsedID {
			return event, true
		}
	}
	return models.Event{}, false
}

func RegisterUserForEvent(reg models.EventRegistration) error {
	user, err := GetUserByID(reg.UserID.String())
	if err != nil {
		return errors.New("user not found")
	}

	for i, event := range eventList {
		if event.ID == reg.EventID {
			for _, participant := range event.Participants {
				if participant.UserID == user.ID {
					return errors.New("user is already registered for this event")
				}
			}

			event.Participants = append(event.Participants, models.Participant{
				UserID: user.ID,
				Role:   reg.Role,
			})

			eventList[i] = event
			return nil
		}
	}

	return errors.New("event not found")
}
