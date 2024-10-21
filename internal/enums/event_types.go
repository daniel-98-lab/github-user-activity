package enums

import "errors"

type EventTypes string

const (
	PushEvent                     EventTypes = "PushEvent"
	CreateEvent                   EventTypes = "CreateEvent"
	DeleteEvent                   EventTypes = "DeleteEvent"
	ForkEvent                     EventTypes = "ForkEvent"
	IssueCommentEvent             EventTypes = "IssueCommentEvent"
	IssuesEvent                   EventTypes = "IssuesEvent"
	PullRequestEvent              EventTypes = "PullRequestEvent"
	PullRequestReviewEvent        EventTypes = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventTypes = "PullRequestReviewCommentEvent"
	WatchEvent                    EventTypes = "WatchEvent"
	ReleaseEvent                  EventTypes = "ReleaseEvent"
)

func ValidateAndAssignEventType(eventType string) (EventTypes, error) {

	// Create a map for the EventTypes constants
	eventTypesMap := map[string]EventTypes{
		string(PushEvent):                     PushEvent,
		string(CreateEvent):                   CreateEvent,
		string(DeleteEvent):                   DeleteEvent,
		string(ForkEvent):                     ForkEvent,
		string(IssueCommentEvent):             IssueCommentEvent,
		string(IssuesEvent):                   IssuesEvent,
		string(PullRequestEvent):              PullRequestEvent,
		string(PullRequestReviewEvent):        PullRequestReviewEvent,
		string(PullRequestReviewCommentEvent): PullRequestReviewCommentEvent,
		string(WatchEvent):                    WatchEvent,
		string(ReleaseEvent):                  ReleaseEvent,
	}

	// Verify if the eventType is in the map
	if event, exists := eventTypesMap[eventType]; exists {
		return event, nil
	}
	// Return error if not exist
	return "", errors.New("event type is not valid")
}
