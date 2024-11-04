package models

import (
	"sync"
)

type ParkingLot struct {
	Capacity       int
	Available      int
	vehicles       map[int]*Vehicle
	spotMutex      sync.Mutex
	AccessGate     sync.Mutex
	OnVehicleEnter func()
	OnVehicleExit  func()
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:  capacity,
		Available: capacity,
		vehicles:  make(map[int]*Vehicle),
	}
}

func (p *ParkingLot) TryEnter(vehicle *Vehicle) bool {
	p.spotMutex.Lock()
	defer p.spotMutex.Unlock()

	if p.Available > 0 {
		p.vehicles[vehicle.ID] = vehicle
		p.Available--
		if p.OnVehicleEnter != nil {
			p.OnVehicleEnter()
		}
		return true
	}
	return false
}

func (p *ParkingLot) Exit(vehicleID int) {
	p.spotMutex.Lock()
	defer p.spotMutex.Unlock()

	if _, exists := p.vehicles[vehicleID]; exists {
		delete(p.vehicles, vehicleID)
		p.Available++
		if p.OnVehicleExit != nil {
			p.OnVehicleExit()
		}
	}
}

func (p *ParkingLot) AccessGateEnter(id int) {
	p.AccessGate.Lock()
}

func (p *ParkingLot) AccessGateExit(id int) {
	p.AccessGate.Unlock()
}
