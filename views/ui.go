package views

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "image/color"
    "parking_simulator/models"
    "sync"
)

var vehicleImages []*canvas.Image

const (
    parkingWidth  = 500
    parkingHeight = 300
    spotWidth     = 60
    spotHeight    = 100
)

func StartUI(parking *models.ParkingLot) {
    a := app.New()
    w := a.NewWindow("Simulador de Estacionamiento")
    w.Resize(fyne.NewSize(600, 400))

    parkingSpaces := container.NewGridWithColumns(5)
    for i := 0; i < parking.Capacity; i++ {
        space := canvas.NewRectangle(color.NRGBA{G: 255, A: 128})
        space.SetMinSize(fyne.NewSize(spotWidth, spotHeight))
        parkingSpaces.Add(space)
    }

    parkingOutline := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
    parkingOutline.SetMinSize(fyne.NewSize(parkingWidth, parkingHeight))
    
    parkingContainer := container.NewMax(parkingOutline, container.NewCenter(parkingSpaces))

    w.SetContent(container.NewCenter(parkingContainer))
    w.Resize(fyne.NewSize(parkingWidth+50, parkingHeight+100))

    updateCh := make(chan int, 20)
    var mu sync.Mutex

    go func() {
        for {
            select {
            case <-updateCh:
                mu.Lock()
                for i, space := range parkingSpaces.Objects {
                    if i < parking.Capacity-parking.Available {
                        space.(*canvas.Rectangle).FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 128}
                    } else {
                        space.(*canvas.Rectangle).FillColor = color.NRGBA{G: 255, A: 128}
                    }
                }
                canvas.Refresh(parkingSpaces)
                mu.Unlock()
            }
        }
    }()

    parking.OnVehicleEnter = func() {
        updateCh <- 1
    }
    parking.OnVehicleExit = func() {
        updateCh <- 1
    }

    w.ShowAndRun()
}
