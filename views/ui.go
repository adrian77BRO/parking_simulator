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

    // Crear el contenedor para los espacios de estacionamiento
    parkingSpaces := container.NewGridWithColumns(5)
    for i := 0; i < parking.Capacity; i++ {
        space := canvas.NewRectangle(color.NRGBA{G: 255, A: 128})
        space.SetMinSize(fyne.NewSize(spotWidth, spotHeight))
        parkingSpaces.Add(space)
    }

    // Crear un contorno alrededor de los espacios de estacionamiento
    parkingOutline := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 255}) // Color blanco
    parkingOutline.SetMinSize(fyne.NewSize(parkingWidth, parkingHeight))
    
    // Contenedor para centrar los espacios de estacionamiento y el contorno
    parkingContainer := container.NewMax(parkingOutline, container.NewCenter(parkingSpaces))

    // Establecer el contenido de la ventana
    w.SetContent(container.NewCenter(parkingContainer)) // Centrar todo en la ventana
    w.Resize(fyne.NewSize(parkingWidth+50, parkingHeight+100)) // Ajuste de tamaño de ventana

    // Canal para actualizar la GUI cuando los vehículos entren o salgan
    updateCh := make(chan int, 20)
    var mu sync.Mutex

    // Función para actualizar la representación gráfica de los vehículos
    go func() {
        for {
            select {
            case <-updateCh:
                mu.Lock()
                // Actualizar representación de cada vehículo en el estacionamiento
                for i, space := range parkingSpaces.Objects {
                    if i < parking.Capacity-parking.Available {
                        // Lugar ocupado
                        space.(*canvas.Rectangle).FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 128}
                    } else {
                        // Lugar libre
                        space.(*canvas.Rectangle).FillColor = color.NRGBA{G: 255, A: 128}
                    }
                }
                canvas.Refresh(parkingSpaces)
                mu.Unlock()
            }
        }
    }()

    // Actualizar GUI en tiempo real cuando cambie el estado del estacionamiento
    parking.OnVehicleEnter = func() {
        updateCh <- 1
    }
    parking.OnVehicleExit = func() {
        updateCh <- 1
    }

    // Mostrar ventana
    w.ShowAndRun()
}
