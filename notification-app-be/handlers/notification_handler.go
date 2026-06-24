package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/services"
	"github.com/affordmedtest/Campus-Evaluation-BE/notification-app-be/utils"

	"github.com/gin-gonic/gin"
)

func GetPriorityInboxHandler(context *gin.Context) {
	topParameter := context.Query("top")
	topN := 10
	if topParameter != "" {
		parsed, err := strconv.Atoi(topParameter)
		if err == nil && parsed > 0 {
			topN = parsed
		}
	}

	log.Printf("Incoming request URL: %s", context.Request.URL.String())
	log.Printf("Requested top count: %d", topN)

	authorizationHeader := context.GetHeader("Authorization")
	tokenDetected := authorizationHeader != ""
	log.Printf("Authorization token detected: %v", tokenDetected)

	if !tokenDetected {
		utils.Log("backend", "warn", "handler", "Missing Authorization header for priority inbox")
		context.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
		return
	}

	bearerToken := ""
	if len(authorizationHeader) > 7 && authorizationHeader[:7] == "Bearer " {
		bearerToken = authorizationHeader[7:]
	} else {
		utils.Log("backend", "warn", "handler", "Invalid Authorization header format for priority inbox")
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
		return
	}

	result, err := services.GetPriorityInbox(topN, bearerToken)
	if err != nil {
		utils.Log("backend", "error", "handler", "Priority inbox fetch failed: "+err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.Log("backend", "info", "handler", "Fetched priority inbox with top="+strconv.Itoa(topN)+" returning "+strconv.Itoa(len(result))+" items")
	context.JSON(http.StatusOK, result)
}
