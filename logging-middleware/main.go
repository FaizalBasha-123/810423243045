package main

import (
	"log"
	"os"

	"github.com/affordmedtest/Campus-Evaluation-BE/logging-middleware/handlers"
	"github.com/affordmedtest/Campus-Evaluation-BE/logging-middleware/middleware"
	"github.com/affordmedtest/Campus-Evaluation-BE/logging-middleware/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Printf("Warning: .env not loaded: %v", err)
		}
	}

	if err := utils.AuthenticateClient(); err != nil {
		log.Fatalf("Auth failed: %v", err)
	}
	log.Printf("Login created successfully")

	utils.Log("backend", "info", "service", "Service starting up")

	router := gin.Default()

	router.Use(middleware.Logger())

	router.GET("/health", handlers.HealthCheck)

	port := os.Getenv("LOGGING_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Logging middleware service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
