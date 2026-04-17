package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// AppState holds shared state for the application
type AppState struct {
	App    fyne.App
	Window fyne.Window
}

// Run boots the Fyne GUI application
func Run() {
	a := app.New()
	w := a.NewWindow("tonet")

	// Set theme (we'll implement this in theme.go later, for now we can just use default)
	// a.Settings().SetTheme(&customTheme{})

	state := &AppState{
		App:    a,
		Window: w,
	}

	// Create tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Send", makeSendTab(state)),
		container.NewTabItem("Receive", makeReceiveTab(state)),
		container.NewTabItem("Relay", makeRelayTab(state)),
		container.NewTabItem("Settings", makeSettingsTab(state)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 400))
	
	// App runs blocking until window is closed
	w.ShowAndRun()
}
