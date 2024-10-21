package enums

type EventTypes string

const (
	PushEvent EventTypes = "PushEvent"
	CreateEvent EventTypes = "CreateEvent"
	DeleteEvent EventTypes = "DeleteEvent"
	ForkEvent EventTypes = "ForkEvent"
	IssueCommentEvent EventTypes = "IssueCommentEvent"
	IssuesEvent EventTypes = "IssuesEvent"
	PullRequestEvent EventTypes = "PullRequestEvent"
	PullRequestReviewEvent EventTypes = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventTypes = "PullRequestReviewCommentEvent"
	WatchEvent EventTypes = "WatchEvent"
	ReleaseEvent EventTypes = "ReleaseEvent"
)