package utils

import (
    "math/rand"
    "time"
)

// RandomArrivalTime genera un tiempo aleatorio para la llegada de vehículos (distribución exponencial)
func RandomArrivalTime() time.Duration {
    rand.Seed(time.Now().UnixNano())
    lambda := 1.0 // Parámetro de tasa para la distribución exponencial
    arrivalInterval := rand.ExpFloat64() / lambda
    return time.Duration(arrivalInterval * float64(time.Second))
}

// RandomParkingDuration genera un tiempo aleatorio entre 3 y 5 segundos para el tiempo de estacionamiento
func RandomParkingDuration() time.Duration {
    rand.Seed(time.Now().UnixNano())
    return time.Duration(3+rand.Intn(3)) * time.Second
}
