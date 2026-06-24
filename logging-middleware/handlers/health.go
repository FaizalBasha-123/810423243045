package handlers

import (
	"net/http"

	"github.com/affordmedtest/Campus-Evaluation-BE/logging-middleware/utils"

	"github.com/gin-gonic/gin"
)

func HealthCheck(context *gin.Context) {
	utils.Log("backend", "info", "handler", "Health check endpoint accessed")
	context.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
