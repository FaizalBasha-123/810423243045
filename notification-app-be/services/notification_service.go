package services

import (
	"fmt"
	"sort"
	"time"
)

type Notification struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func GetPriorityInbox(topN int, bearerToken string) ([]Notification, error) {
	if bearerToken == "" {
		return nil, fmt.Errorf("missing bearer token")
	}

	sampleNotifications := []Notification{
		{ID: "n1", Type: "Event", Message: "System maintenance scheduled", Timestamp: "2026-04-22 10:00:00"},
		{ID: "n2", Type: "Result", Message: "Lab results published", Timestamp: "2026-04-22 11:30:00"},
		{ID: "n3", Type: "Placement", Message: "Placement drive registration open", Timestamp: "2026-04-22 09:00:00"},
		{ID: "n4", Type: "Result", Message: "Semester results announced", Timestamp: "2026-04-22 10:45:00"},
		{ID: "n5", Type: "Event", Message: "Tech fest registration deadline", Timestamp: "2026-04-22 08:00:00"},
		{ID: "n6", Type: "Event", Message: "Workshop on AI and ML", Timestamp: "2026-04-22 12:00:00"},
		{ID: "n7", Type: "Placement", Message: "Mock interview schedule", Timestamp: "2026-04-22 11:00:00"},
		{ID: "n8", Type: "Result", Message: "Quiz competition results", Timestamp: "2026-04-22 09:15:00"},
		{ID: "n9", Type: "Result", Message: "Internship selection list", Timestamp: "2026-04-22 14:00:00"},
		{ID: "n10", Type: "Placement", Message: "Company onboarding details", Timestamp: "2026-04-22 07:30:00"},
		{ID: "n11", Type: "Event", Message: "Cultural fest registrations open", Timestamp: "2026-04-22 13:00:00"},
		{ID: "n12", Type: "Event", Message: "Sports meet schedule", Timestamp: "2026-04-22 09:45:00"},
		{ID: "n13", Type: "Placement", Message: "Resume submission deadline extended", Timestamp: "2026-04-22 10:30:00"},
		{ID: "n14", Type: "Placement", Message: "Pre-placement talk scheduled", Timestamp: "2026-04-22 08:45:00"},
		{ID: "n15", Type: "Event", Message: "Hackathon team formation", Timestamp: "2026-04-22 11:15:00"},
	}

	weightMap := map[string]int{
		"Placement": 3,
		"Result":    2,
		"Event":     1,
	}

	sort.Slice(sampleNotifications, func(i, j int) bool {
		weightI := weightMap[sampleNotifications[i].Type]
		weightJ := weightMap[sampleNotifications[j].Type]
		if weightI != weightJ {
			return weightI > weightJ
		}

		timeI, errI := time.Parse("2006-01-02 15:04:05", sampleNotifications[i].Timestamp)
		timeJ, errJ := time.Parse("2006-01-02 15:04:05", sampleNotifications[j].Timestamp)
		if errI == nil && errJ == nil {
			return timeI.After(timeJ)
		}
		return sampleNotifications[i].Timestamp > sampleNotifications[j].Timestamp
	})

	if topN <= 0 {
		return nil, fmt.Errorf("topN must be positive")
	}

	if topN > len(sampleNotifications) {
		return sampleNotifications, nil
	}

	return sampleNotifications[:topN], nil
}
