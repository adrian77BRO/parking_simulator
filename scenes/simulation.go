package scenes

import (
	"fmt"
	"parking_simulator/models"
	"parking_simulator/utils"
	"time"
)

func Simulation(parking *models.ParkingLot) {
	for vehicleID := 1; vehicleID <= 100; vehicleID++ {
		time.Sleep(utils.RandomArrivalTime())

		parkingTime := utils.RandomParkingDuration()
		vehicle := models.NewVehicle(vehicleID, parkingTime)

		go handleVehicle(parking, vehicle)
	}
}

func handleVehicle(parking *models.ParkingLot, vehicle *models.Vehicle) {
	fmt.Printf("Vehículo %d ha llegado.\n", vehicle.ID)

	if parking.TryEnter(vehicle) {
		parking.AccessGateEnter(vehicle.ID)
		fmt.Printf("Vehículo %d entra al estacionamiento.\n", vehicle.ID)
		parking.AccessGateExit(vehicle.ID)

		time.Sleep(vehicle.ParkingTime)

		parking.AccessGateEnter(vehicle.ID)
		fmt.Printf("Vehículo %d abandonando el estacionamiento.\n", vehicle.ID)
		parking.Exit(vehicle.ID)
		parking.AccessGateExit(vehicle.ID)

		fmt.Printf("Vehículo %d ha salido del estacionamiento.\n", vehicle.ID)
	} else {
		fmt.Printf("Vehículo %d en espera de un lugar disponible.\n", vehicle.ID)
	}
}
