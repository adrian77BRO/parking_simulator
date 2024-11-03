package models

import "sync"

type ParkingLot struct {
    Capacity   int
    Available  int
    AccessGate sync.Mutex
    SpotMutex  sync.Mutex

    // Callbacks para actualizar la GUI
    OnVehicleEnter func()
    OnVehicleExit  func()
}

// Crear un nuevo estacionamiento
func NewParkingLot(capacity int) *ParkingLot {
    return &ParkingLot{
        Capacity:  capacity,
        Available: capacity,
    }
}

// Intentar entrar si hay lugar
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

// Salir, liberando un espacio
func (p *ParkingLot) Exit(id int) {
    p.SpotMutex.Lock()
    defer p.SpotMutex.Unlock()

    p.Available++
    if p.OnVehicleExit != nil {
        p.OnVehicleExit()
    }
}

// AccessGateEnter simula que un vehículo accede a la entrada del estacionamiento
func (p *ParkingLot) AccessGateEnter(id int) {
    p.AccessGate.Lock()
}

// AccessGateExit simula que un vehículo libera la entrada del estacionamiento
func (p *ParkingLot) AccessGateExit(id int) {
    p.AccessGate.Unlock()
}
