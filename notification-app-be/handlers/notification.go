package handlers

import (
	"fmt"
	"net/http"

	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/models"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/services"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/utils"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	Service *services.NotificationService
}

func NewNotificationHandler(service *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{Service: service}
}

func (handler *NotificationHandler) CreateNotification(context *gin.Context) {
	var request models.Notification
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.UserID == "" || request.Title == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user_id and title are required"})
		return
	}

	notification := handler.Service.Create(request)
	utils.Log("backend", "info", "handler", "Created notification "+notification.ID+" for user "+request.UserID)
	context.JSON(http.StatusCreated, notification)
}

func (handler *NotificationHandler) GetNotification(context *gin.Context) {
	id := context.Param("id")
	notification, exists := handler.Service.GetByID(id)
	if !exists {
		utils.Log("backend", "warn", "handler", "Notification "+id+" not found")
		context.JSON(http.StatusNotFound, gin.H{"error": "notification not found"})
		return
	}
	utils.Log("backend", "info", "handler", "Retrieved notification "+id)
	context.JSON(http.StatusOK, notification)
}

func (handler *NotificationHandler) GetUserNotifications(context *gin.Context) {
	userID := context.Param("userID")
	notifications := handler.Service.GetByUserID(userID)
	utils.Log("backend", "info", "handler", "Listed "+fmt.Sprintf("%d", len(notifications))+" notifications for user "+userID)
	context.JSON(http.StatusOK, notifications)
}
