package handlers

import (
	"fmt"
	"net/http"

	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/models"
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/services"
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/utils"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	Service *services.ScheduleService
}

func NewScheduleHandler(service *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{Service: service}
}

func (handler *ScheduleHandler) CreateVehicle(context *gin.Context) {
	var request models.Vehicle
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" || request.Type == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "name and type are required"})
		return
	}

	vehicle := handler.Service.CreateVehicle(request)
	utils.Log("backend", "info", "handler", "Created vehicle "+vehicle.ID)
	context.JSON(http.StatusCreated, vehicle)
}

func (handler *ScheduleHandler) ListVehicles(context *gin.Context) {
	vehicles := handler.Service.ListVehicles()
	utils.Log("backend", "info", "handler", "Listed "+fmt.Sprintf("%d", len(vehicles))+" vehicles")
	context.JSON(http.StatusOK, vehicles)
}

func (handler *ScheduleHandler) CreateSchedule(context *gin.Context) {
	var request models.Schedule
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.VehicleID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "vehicle_id is required"})
		return
	}

	schedule, err := handler.Service.CreateSchedule(request)
	if err != nil {
		utils.Log("backend", "warn", "handler", "Schedule creation failed: "+err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.Log("backend", "info", "handler", "Created schedule "+schedule.ID+" for vehicle "+request.VehicleID)
	context.JSON(http.StatusCreated, schedule)
}
