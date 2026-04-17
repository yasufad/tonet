package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeSettingsTab(state *AppState) fyne.CanvasObject {
	// UI Elements
	relayEntry := widget.NewEntry()
	relayEntry.SetPlaceHolder("Custom relay address (e.g. relay.example.com:9009)")

	relayIPv6Entry := widget.NewEntry()
	relayIPv6Entry.SetPlaceHolder("Custom IPv6 relay address")

	relayPassEntry := widget.NewPasswordEntry()
	relayPassEntry.SetPlaceHolder("Relay password (if required)")

	curveSelect := widget.NewSelect([]string{"p256", "p384", "p521", "siec"}, nil)
	curveSelect.SetSelected("p256") // default

	socks5Entry := widget.NewEntry()
	socks5Entry.SetPlaceHolder("SOCKS5 proxy (e.g. 127.0.0.1:9050)")

	httpProxyEntry := widget.NewEntry()
	httpProxyEntry.SetPlaceHolder("HTTP proxy (e.g. 127.0.0.1:8080)")

	throttleEntry := widget.NewEntry()
	throttleEntry.SetPlaceHolder("Upload throttle (e.g. 500k)")

	classicModeCheck := widget.NewCheck("Classic mode (CROC_SECRET environmental behaviour)", nil)

	saveBtn := widget.NewButton("Save Settings", func() {
		// Placeholder for saving settings logic (writes to config dir)
	})
	saveBtn.Importance = widget.HighImportance

	// Form layout
	form := widget.NewForm(
		widget.NewFormItem("Relay Address", relayEntry),
		widget.NewFormItem("Relay IPv6", relayIPv6Entry),
		widget.NewFormItem("Relay Password", relayPassEntry),
		widget.NewFormItem("Encryption Curve", curveSelect),
		widget.NewFormItem("SOCKS5 Proxy", socks5Entry),
		widget.NewFormItem("HTTP Proxy", httpProxyEntry),
		widget.NewFormItem("Upload Throttle", throttleEntry),
	)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Application Settings", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		form,
		widget.NewSeparator(),
		classicModeCheck,
		widget.NewSeparator(),
		saveBtn,
	)

	// Wrap in a scroll container in case it exceeds window height
	return container.NewPadded(container.NewVScroll(content))
}
