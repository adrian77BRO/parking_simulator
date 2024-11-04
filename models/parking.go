package models

import "sync"

type ParkingLot struct {
    Capacity   int
    Available  int
    AccessGate sync.Mutex
    SpotMutex  sync.Mutex

    OnVehicleEnter func()
    OnVehicleExit  func()
}

func NewParkingLot(capacity int) *ParkingLot {
    return &ParkingLot{
        Capacity:  capacity,
        Available: capacity,
    }
}

func (p *ParkingLot) TryEnter(id int) bool {
    p.SpotMutex.Lock()
    defer p.SpotMutex.Unlock()

    if p.Available > 0 {
        p.Available--
        if p.OnVehicleEnter != nil {
            p.OnVehicleEnter()
        }
        return true
    }
    return false
}

func (p *ParkingLot) Exit(id int) {
    p.SpotMutex.Lock()
    defer p.SpotMutex.Unlock()

    p.Available++
    if p.OnVehicleExit != nil {
        p.OnVehicleExit()
    }
}

func (p *ParkingLot) AccessGateEnter(id int) {
    p.AccessGate.Lock()
}

func (p *ParkingLot) AccessGateExit(id int) {
    p.AccessGate.Unlock()
}
