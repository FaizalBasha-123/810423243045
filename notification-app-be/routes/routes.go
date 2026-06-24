package routes

import (
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	httpRouter := gin.Default()

	v1 := httpRouter.Group("/api/v1")
	v1.GET("/priority-inbox", handlers.GetPriorityInboxHandler)

	return httpRouter
}
