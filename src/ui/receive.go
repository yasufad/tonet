package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeReceiveTab(state *AppState) fyne.CanvasObject {
	label := widget.NewLabel("Receive Files")
	// Placeholder for receive UI
	return container.NewVBox(label)
}
