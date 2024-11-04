package models

import "time"

type Vehicle struct {
	ID          int
	ParkingTime time.Duration
}

func NewVehicle(id int, parkingTime time.Duration) *Vehicle {
	return &Vehicle{
		ID:          id,
		ParkingTime: parkingTime,
	}
}
