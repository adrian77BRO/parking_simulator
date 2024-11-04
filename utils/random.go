package utils

import (
	"math/rand"
	"time"
)

func RandomArrivalTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	lambda := 1.0
	arrivalInterval := rand.ExpFloat64() / lambda
	return time.Duration(arrivalInterval * float64(time.Second))
}

func RandomParkingDuration() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(3+rand.Intn(3)) * time.Second
}
