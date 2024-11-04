package models 

type Vehicle struct {
    ID int
    Done chan bool
}

func NewVehicle(id int) *Vehicle {
    return &Vehicle{
        ID:   id,
        Done: make(chan bool),
    }
}
