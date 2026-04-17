package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeSettingsTab(state *AppState) fyne.CanvasObject {
	label := widget.NewLabel("Settings")
	// Placeholder for settings UI
	return container.NewVBox(label)
}
