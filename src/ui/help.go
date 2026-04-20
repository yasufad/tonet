package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeHelpTab(state *AppState) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabelWithStyle("Tonet - Secure File Transfer", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("How to Send Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("1. Go to the Send tab"),
		widget.NewLabel("2. Click 'Select Files' or 'Select Folder'"),
		widget.NewLabel("3. Choose your hash algorithm and options"),
		widget.NewLabel("4. Click 'Send'"),
		widget.NewLabel("5. Share the generated code with the recipient"),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("How to Receive Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("1. Go to the Receive tab"),
		widget.NewLabel("2. Enter the code phrase from the sender"),
		widget.NewLabel("3. Optionally select an output directory"),
		widget.NewLabel("4. Click 'Receive'"),
		widget.NewLabel("5. Files will be downloaded to the selected location"),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("Running Your Own Relay", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("1. Go to the Relay tab"),
		widget.NewLabel("2. Configure host (optional) and base port"),
		widget.NewLabel("3. Set number of transfer ports (minimum 2)"),
		widget.NewLabel("4. Click 'Start Relay'"),
		widget.NewLabel("5. Update Settings tab with your relay address"),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("Settings", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Configure relay servers, encryption curves, and proxy settings."),
		widget.NewLabel("Settings are saved and used by Send and Receive operations."),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("Security", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("• End-to-end encrypted using PAKE"),
		widget.NewLabel("• Code phrases are used for secure key exchange"),
		widget.NewLabel("• Files are encrypted before transmission"),
		widget.NewLabel("• Supports custom relay servers for privacy"),
		widget.NewSeparator(),

		widget.NewLabelWithStyle("CLI Mode", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("tonet also works from the command line:"),
		widget.NewLabel("  tonet send file.txt"),
		widget.NewLabel("  tonet <code-phrase>"),
		widget.NewSeparator(),

		widget.NewLabel("For more information, visit: github.com/yasufad/tonet"),
	)

	return container.NewPadded(container.NewVScroll(content))
}
