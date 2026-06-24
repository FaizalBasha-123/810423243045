package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/models"
	"github.com/affordmedtest/Campus-Evaluation-BE/vehicle-scheduler-be/utils"
)

type ScheduleService struct {
	mu          sync.Mutex
	vehicles    map[string]models.Vehicle
	schedules   map[string]models.Schedule
	vehicleSeq  int64
	scheduleSeq int64
}

func NewScheduleService() *ScheduleService {
	return &ScheduleService{
		vehicles:    make(map[string]models.Vehicle),
		schedules:   make(map[string]models.Schedule),
		vehicleSeq:  1,
		scheduleSeq: 1,
	}
}

func (service *ScheduleService) CreateVehicle(request models.Vehicle) models.Vehicle {
	service.mu.Lock()
	defer service.mu.Unlock()

	id := fmt.Sprintf("v%d", service.vehicleSeq)
	service.vehicleSeq++

	vehicle := models.Vehicle{
		ID:     id,
		Name:   request.Name,
		Type:   request.Type,
		Status: "available",
	}

	service.vehicles[id] = vehicle
	utils.Log("backend", "info", "service", "Stored vehicle "+id+" in memory")
	return vehicle
}

func (service *ScheduleService) ListVehicles() []models.Vehicle {
	service.mu.Lock()
	defer service.mu.Unlock()

	var result []models.Vehicle
	for _, vehicle := range service.vehicles {
		result = append(result, vehicle)
	}
	return result
}

func (service *ScheduleService) CreateSchedule(request models.Schedule) (models.Schedule, error) {
	service.mu.Lock()
	defer service.mu.Unlock()

	if _, exists := service.vehicles[request.VehicleID]; !exists {
		return models.Schedule{}, fmt.Errorf("vehicle with id %s not found", request.VehicleID)
	}

	id := fmt.Sprintf("s%d", service.scheduleSeq)
	service.scheduleSeq++

	schedule := models.Schedule{
		ID:          id,
		VehicleID:   request.VehicleID,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		Description: request.Description,
	}

	if schedule.StartTime.IsZero() {
		schedule.StartTime = time.Now()
	}
	if schedule.EndTime.IsZero() {
		schedule.EndTime = schedule.StartTime.Add(time.Hour)
	}

	service.schedules[id] = schedule
	utils.Log("backend", "info", "service", "Stored schedule "+id+" in memory")
	return schedule, nil
}
