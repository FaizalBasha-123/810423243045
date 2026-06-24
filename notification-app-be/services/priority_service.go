package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/models"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/utils"
)

func GetPriorityInbox(topN int, bearerToken string) ([]models.Notification, error) {
	if bearerToken == "" {
		return nil, fmt.Errorf("missing bearer token")
	}

	baseURL := "http://4.224.186.213"
	request, err := http.NewRequest("GET", baseURL+"/evaluation-service/priority-inbox", nil)
	if err != nil {
		utils.Log("backend", "error", "service", "Failed to create priority inbox request: "+err.Error())
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		utils.Log("backend", "error", "service", "Priority inbox request failed: "+err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		utils.Log("backend", "error", "service", "Priority inbox returned status "+fmt.Sprintf("%d", response.StatusCode))
		return nil, fmt.Errorf("external service returned status %d", response.StatusCode)
	}

	var result []models.Notification
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		utils.Log("backend", "error", "service", "Failed to decode priority inbox response: "+err.Error())
		return nil, err
	}

	if len(result) > topN {
		result = result[:topN]
	}

	return result, nil
}
