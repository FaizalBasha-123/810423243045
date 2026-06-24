package handlers

import (
	"net/http"
	"time"

	"github.com/affordmedtest/Campus-Evaluation-BE/logging-middleware/utils"

	"github.com/gin-gonic/gin"
)

func HealthCheck(context *gin.Context) {
	utils.Log("backend", "info", "handler", "Manual health check triggered by Postman")

	context.JSON(http.StatusOK, gin.H{
		"status":    "UP",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
