package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

type Notification struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type priorityInboxResponse struct {
	Notifications []Notification `json:"notifications"`
}

func GetPriorityInbox(topN int, bearerToken string) ([]Notification, error) {
	if bearerToken == "" {
		return nil, fmt.Errorf("missing bearer token")
	}

	request, err := http.NewRequest("GET", "http://4.224.186.213/evaluation-service/priority-inbox", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external service returned status %d", response.StatusCode)
	}

	var wrapper priorityInboxResponse
	if err := json.NewDecoder(response.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	weightMap := map[string]int{
		"Placement": 3,
		"Result":    2,
		"Event":     1,
	}

	sort.Slice(wrapper.Notifications, func(i, j int) bool {
		weightI := weightMap[wrapper.Notifications[i].Type]
		weightJ := weightMap[wrapper.Notifications[j].Type]
		if weightI != weightJ {
			return weightI > weightJ
		}

		timeI, errI := time.Parse("2006-01-02 15:04:05", wrapper.Notifications[i].Timestamp)
		timeJ, errJ := time.Parse("2006-01-02 15:04:05", wrapper.Notifications[j].Timestamp)
		if errI == nil && errJ == nil {
			return timeI.After(timeJ)
		}
		return wrapper.Notifications[i].Timestamp > wrapper.Notifications[j].Timestamp
	})

	if topN <= 0 {
		return nil, fmt.Errorf("topN must be positive")
	}

	if topN > len(wrapper.Notifications) {
		return wrapper.Notifications, nil
	}

	return wrapper.Notifications[:topN], nil
}
