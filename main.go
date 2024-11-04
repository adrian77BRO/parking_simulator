package main

import (
    "parking_simulator/models"
    "parking_simulator/scenes"
    "parking_simulator/views"
)

func main() {
    parking := models.NewParkingLot(20)

    go scenes.Simulation(parking)

    views.StartUI(parking)
}
