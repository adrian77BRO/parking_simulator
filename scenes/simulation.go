package scenes

import (
    "fmt"
    "parking_simulator/models"
    "parking_simulator/utils"
    "time"
)

// Simulation ejecuta la simulación de vehículos en el estacionamiento
func Simulation(parking *models.ParkingLot) {
    vehicleID := 1

    // Genera vehículos en un ciclo infinito
    for {
        // Intervalo de llegada de vehículos usando distribución de Poisson
        time.Sleep(utils.RandomArrivalTime())
        
        // Crear vehículo y lanzar una goroutine para simular su flujo de entrada y salida
        go func(id int) {
            fmt.Printf("Vehicle %d arrived.\n", id)

            // Intentar entrar al estacionamiento
            if parking.TryEnter(id) {
                parking.AccessGateEnter(id)
                fmt.Printf("Vehicle %d accessed the gate and entered the parking.\n", id)
                parking.AccessGateExit(id)

                // Tiempo que el vehículo estará estacionado
                time.Sleep(utils.RandomParkingDuration())

                // Simular salida del vehículo
                parking.AccessGateEnter(id)
                fmt.Printf("Vehicle %d is leaving the parking.\n", id)
                parking.Exit(id)
                parking.AccessGateExit(id)

                fmt.Printf("Vehicle %d has exited the parking.\n", id)
            } else {
                fmt.Printf("Vehicle %d is waiting for an available spot.\n", id)
            }
        }(vehicleID)

        vehicleID++
    }
}
