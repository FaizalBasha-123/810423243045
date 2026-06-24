package main

import (
	"log"
	"os"

	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/routes"
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/utils"

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

	router := routes.SetupRouter()

	port := os.Getenv("VEHICLE_SCHEDULER_PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Vehicle scheduler service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
