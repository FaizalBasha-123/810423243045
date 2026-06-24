package models

import "time"

type Vehicle struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type Schedule struct {
	ID          string    `json:"id"`
	VehicleID   string    `json:"vehicle_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}
