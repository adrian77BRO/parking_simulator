package models 

// Vehicle representa un vehículo en el estacionamiento.
type Vehicle struct {
    ID int          // ID único del vehículo
    Done chan bool  // Canal que indicará cuando el vehículo haya terminado su operación
}

// NewVehicle crea un nuevo vehículo con un ID y un canal para la terminación.
func NewVehicle(id int) *Vehicle {
    return &Vehicle{
        ID:   id,
        Done: make(chan bool),
    }
}
