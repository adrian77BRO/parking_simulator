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
        
        go func(id int) {
            fmt.Printf("Vehículo %d ha llegado.\n", id)

            if parking.TryEnter(id) {
                parking.AccessGateEnter(id)
                fmt.Printf("Vehículo %d accede a la puerta y entra al estacionamiento.\n", id)
                parking.AccessGateExit(id)

                time.Sleep(utils.RandomParkingDuration())

                parking.AccessGateEnter(id)
                fmt.Printf("Vehículo %d abandonando el estacionamiento.\n", id)
                parking.Exit(id)
                parking.AccessGateExit(id)

                fmt.Printf("Vehículo %d ha salido del estacionamiento.\n", id)
            } else {
                fmt.Printf("Vehículo %d en espera de un lugar disponible.\n", id)
            }
        }(vehicleID)
    }
}
