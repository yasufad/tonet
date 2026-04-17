package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeSendTab(state *AppState) fyne.CanvasObject {
	label := widget.NewLabel("Send Files")
	// Placeholder for send UI
	return container.NewVBox(label)
}
