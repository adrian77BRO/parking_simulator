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

const (
	parkingWidth  = 600
	parkingHeight = 480
	spotSize      = 100
)

func StartUI(parking *models.ParkingLot) {
	a := app.New()
	w := a.NewWindow("Simulador de estacionamiento")
	w.Resize(fyne.NewSize(600, 400))

	parkingSpaces := container.NewGridWithColumns(5)
	spaceContainers := make([]*fyne.Container, parking.Capacity)
	for i := 0; i < parking.Capacity; i++ {
		spaceRect := canvas.NewRectangle(color.NRGBA{G: 255, A: 128})
		spaceRect.SetMinSize(fyne.NewSize(spotSize, spotSize))

		spaceContainer := container.NewMax(spaceRect)
		spaceContainers[i] = spaceContainer
		parkingSpaces.Add(spaceContainer)
	}

	parkingOutline := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	parkingOutline.SetMinSize(fyne.NewSize(parkingWidth, parkingHeight))

	parkingContainer := container.NewMax(parkingOutline, container.NewCenter(parkingSpaces))
	w.SetContent(container.NewCenter(parkingContainer))
	w.Resize(fyne.NewSize(parkingWidth+50, parkingHeight+100))

	updateCh := make(chan int, 20)
	var mu sync.Mutex

	go updateParkingSpaces(spaceContainers, parkingSpaces, parking, updateCh, &mu)

	parking.OnVehicleEnter = func() {
		updateCh <- 1
	}
	parking.OnVehicleExit = func() {
		updateCh <- 1
	}

	w.ShowAndRun()
}

func updateParkingSpaces(spaceContainers []*fyne.Container, parkingSpaces *fyne.Container, parking *models.ParkingLot, updateCh chan int, mu *sync.Mutex) {
	for {
		select {
		case <-updateCh:
			mu.Lock()
			for i, spaceContainer := range spaceContainers {
				spaceContainer.Objects = nil

				if i < parking.Capacity-parking.Available {
					background := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 128})
					background.SetMinSize(fyne.NewSize(spotSize, spotSize))
					spaceContainer.Add(background)

					car := canvas.NewImageFromFile("assets/car.png")
					car.FillMode = canvas.ImageFillContain
					spaceContainer.Add(car)
				} else {
					background := canvas.NewRectangle(color.NRGBA{G: 255, A: 128})
					background.SetMinSize(fyne.NewSize(spotSize, spotSize))
					spaceContainer.Add(background)
				}
			}
			canvas.Refresh(parkingSpaces)
			mu.Unlock()
		}
	}
}
