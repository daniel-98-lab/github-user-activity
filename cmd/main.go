package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/daniel-98-lab/github-user-activity/internal/services"
)

func main() {

    service := services.UserActivity{};

    if len(os.Args) < 1 {
        fmt.Println("No user provided.")
        return
    }

    args := os.Args
	user := strings.Join(args[1:], " ")

	events, err := service.LoadEvents(user);

	if err != nil {
		fmt.Printf("Error loading events: %v\n", err)
		return
	}

	for _, event := range events {
		fmt.Printf("Event ID: %s, Type: %s\n", event.ID, event.Type)
	}
}