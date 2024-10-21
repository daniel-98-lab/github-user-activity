package models

import (
	"time"
	"github.com/daniel-98-lab/github-user-activity/internal/enums"
)

// Actor represents the user who performed the event
type Actor struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID  string `json:"gravatar_id"`
	URL         string `json:"url"`
	AvatarURL   string `json:"avatar_url"`
}

// Repo represents the repository where the event occurred
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// CommitAuthor represents the author of a commit
type CommitAuthor struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Commit represents an individual commit
type Commit struct {
	SHA     string       `json:"sha"`
	Author  CommitAuthor `json:"author"`
	Message string       `json:"message"`
	Distinct bool        `json:"distinct"`
	URL     string       `json:"url"`
}

// Payload represents the data related to the event
type Payload struct {
	RepositoryID  int      `json:"repository_id"`
	PushID        int64    `json:"push_id"`
	Size          int      `json:"size"`
	DistinctSize  int      `json:"distinct_size"`
	Ref           string   `json:"ref"`
	Head          string   `json:"head"`
	Before        string   `json:"before"`
	Commits       []Commit `json:"commits"`
}

// Event represents the GitHub event
type Event struct {
	ID        string    			`json:"id"`
	Type      enums.EventTypes    	`json:"type"`
	Actor     Actor     			`json:"actor"`
	Repo      Repo      			`json:"repo"`
	Payload   Payload   			`json:"payload"`
	Public    bool      			`json:"public"`
	CreatedAt time.Time 			`json:"created_at"`
}
