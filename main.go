package main

import (
    "parking_simulator/models"
    "parking_simulator/scenes"
    "parking_simulator/views"
)

func main() {
    // Crear estacionamiento con una capacidad de 20 vehículos
    parking := models.NewParkingLot(20)

    // Iniciar la simulación del estacionamiento en una goroutine
    go scenes.Simulation(parking)

    // Iniciar la interfaz gráfica (debe estar en la goroutine principal)
    views.StartUI(parking)
}
