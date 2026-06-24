package routes

import (
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/handlers"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	notificationService := services.NewNotificationService()
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	notificationGroup := router.Group("/notifications")
	{
		notificationGroup.POST("", notificationHandler.CreateNotification)
		notificationGroup.GET("/:id", notificationHandler.GetNotification)
		notificationGroup.GET("/user/:userID", notificationHandler.GetUserNotifications)
		notificationGroup.GET("/priority-inbox", notificationHandler.PriorityInbox)
	}

	return router
}
