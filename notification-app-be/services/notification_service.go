package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/models"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/utils"
)

type NotificationService struct {
	mu            sync.Mutex
	notifications map[string]models.Notification
	nextID        int64
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		notifications: make(map[string]models.Notification),
		nextID:        1,
	}
}

func (service *NotificationService) Create(request models.Notification) models.Notification {
	service.mu.Lock()
	defer service.mu.Unlock()

	id := fmt.Sprintf("n%d", service.nextID)
	service.nextID++

	notification := models.Notification{
		ID:        id,
		UserID:    request.UserID,
		Title:     request.Title,
		Message:   request.Message,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	service.notifications[id] = notification
	utils.Log("backend", "info", "service", "Stored notification "+id+" in memory")
	return notification
}

func (service *NotificationService) GetByID(id string) (models.Notification, bool) {
	service.mu.Lock()
	defer service.mu.Unlock()

	notification, exists := service.notifications[id]
	return notification, exists
}

func (service *NotificationService) GetByUserID(userID string) []models.Notification {
	service.mu.Lock()
	defer service.mu.Unlock()

	var result []models.Notification
	for _, notification := range service.notifications {
		if notification.UserID == userID {
			result = append(result, notification)
		}
	}
	return result
}
