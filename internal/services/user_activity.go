package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/daniel-98-lab/github-user-activity/internal/enums"
	"github.com/daniel-98-lab/github-user-activity/internal/models"
)

type UserActivity struct{}

// LoadEvents is a http fetch to give user events
func (s *UserActivity) LoadEvents(user string, eventTypeInput string) ([]models.Event, error) {

	if user == "" {
		return nil, errors.New("no user provided")
	}

	eventsGetURL := fmt.Sprintf("https://api.github.com/users/%s/events", user)
	resp, err := http.Get(eventsGetURL)

	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch events: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var events []models.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// If eventTypeInput is void, return all events
	if eventTypeInput == "" {
		return events, nil
	}

	eventType, err := enums.ValidateAndAssignEventType(eventTypeInput)

	// Validate and Assign EventType
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Event type valid:", eventType)
	}

	// Filter events by the provided type
	var filteredEvents []models.Event
	for _, event := range events {
		if event.Type == enums.EventTypes(eventType) {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents, nil
}
