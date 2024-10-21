package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/daniel-98-lab/github-user-activity/internal/services"
)

func main() {

	service := services.UserActivity{}

	if len(os.Args) < 2 {
		fmt.Println("Usage: user-activity <username> [eventType]")
		return
	}

	args := os.Args
	user := args[1]
	event := strings.Join(args[2:], " ")

	events, err := service.LoadEvents(user, event)

	if err != nil {
		fmt.Printf("Error loading events: %v\n", err)
		return
	}

	for _, event := range events {
		fmt.Printf("Event ID: %s, Type: %s\n", event.ID, event.Type)
	}
}
