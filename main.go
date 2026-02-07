package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type EventType string

const (
	EVENT_TYPE_CREATE EventType = "CreateEvent"
	EVENT_TYPE_PUSH   EventType = "PushEvent"
	EVENT_TYPE_ISSUES EventType = "IssuesEvent"
)

type Event struct{
	// ID int `json:"id"`
	Type EventType `json:"type"`
	Repo struct{
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct{
		Action string `json:"action"`
		RefType string `json:"ref_type"`
	} `json:"payload"`
}

type Events []Event

func (e Event) ProcessPushEvent() {
	fmt.Printf("Pushed to %s\n", e.Repo.Name)
}

func (e Event) ProcessIssuesEvent() {
	switch e.Payload.Action {
	case "opened":
		fmt.Printf("Opened a new issue in %s\n", e.Repo.Name)
	case "closed":
		fmt.Printf("Closed an issue in %s\n", e.Repo.Name)
	case "reopened":
		fmt.Printf("Reopened an issue in %s\n", e.Repo.Name)
	}
}

func (e Event) ProcessCreateEvent() {
	switch e.Payload.RefType {
	case "branch":
		fmt.Printf("Created a branch in %s\n", e.Repo.Name)
	case "tag":
		fmt.Printf("Created a tag in %s\n", e.Repo.Name)
	case "repository":
		fmt.Printf("Created a repository %s\n", e.Repo.Name)
	}
}

func (e Event) ProcessEvent() {
	switch e.Type {
	case EVENT_TYPE_PUSH:
		e.ProcessPushEvent()
	case EVENT_TYPE_ISSUES:
		e.ProcessIssuesEvent()
	case EVENT_TYPE_CREATE:
		e.ProcessCreateEvent()
	default:
		fmt.Printf("Skipping unknown event of type %s\n", e.Type)
	}
}

const url = "https://api.github.com/users/%s/events"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Script takes username as argument")
		os.Exit(1)
	}
	username := os.Args[1]

	r, err := http.Get(fmt.Sprintf(url, username))
	if err != nil {
		log.Fatalf("Error fetching GitHub data: %v", err)
	}
	defer r.Body.Close()

	var events Events
	err = json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad JSON: %s\n", err)
		os.Exit(1)
	}

	for _, event := range events {
		event.ProcessEvent()
	}
}
