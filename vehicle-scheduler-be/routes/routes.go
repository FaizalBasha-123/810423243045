package routes

import (
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/handlers"
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	scheduleService := services.NewScheduleService()
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	vehicleGroup := router.Group("/vehicles")
	{
		vehicleGroup.POST("", scheduleHandler.CreateVehicle)
		vehicleGroup.GET("", scheduleHandler.ListVehicles)
	}

	scheduleGroup := router.Group("/schedules")
	{
		scheduleGroup.POST("", scheduleHandler.CreateSchedule)
	}

	return router
}
