package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeRelayTab(state *AppState) fyne.CanvasObject {
	label := widget.NewLabel("Relay Management")
	// Placeholder for relay UI
	return container.NewVBox(label)
}
