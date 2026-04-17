package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeRelayTab(state *AppState) fyne.CanvasObject {
	// UI Elements
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Optional: host/IP to bind to (default: all interfaces)")

	portsEntry := widget.NewEntry()
	portsEntry.SetText("9009,9010,9011,9012,9013") // default ports
	
	statusLabel := widget.NewLabel("Relay Status: Stopped")
	statusLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Action buttons
	var startBtn *widget.Button
	var stopBtn *widget.Button

	startBtn = widget.NewButton("Start Relay", func() {
		// Placeholder for starting relay
		statusLabel.SetText("Relay Status: Running")
		startBtn.Disable()
		stopBtn.Enable()
	})
	startBtn.Importance = widget.HighImportance

	stopBtn = widget.NewButton("Stop Relay", func() {
		// Placeholder for stopping relay
		statusLabel.SetText("Relay Status: Stopped")
		startBtn.Enable()
		stopBtn.Disable()
	})
	stopBtn.Disable()

	// Connection log
	logData := []string{"Ready to accept connections."}
	logList := widget.NewList(
		func() int {
			return len(logData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template Label")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(logData[i])
		},
	)

	// Layout assembly
	configBox := container.NewVBox(
		widget.NewLabel("Host bind address:"),
		hostEntry,
		widget.NewLabel("Ports (comma-separated):"),
		portsEntry,
	)

	actionBox := container.NewHBox(startBtn, stopBtn)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Relay Management", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		configBox,
		widget.NewSeparator(),
		statusLabel,
		actionBox,
		widget.NewSeparator(),
		widget.NewLabel("Connection Log:"),
	)

	// We use a Border layout to let the logList expand and fill remaining space
	return container.NewPadded(container.NewBorder(content, nil, nil, nil, logList))
}
