package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func makeReceiveTab(state *AppState) fyne.CanvasObject {
	// UI Elements
	codeEntry := widget.NewEntry()
	codeEntry.SetPlaceHolder("Enter receive code (e.g., fast-pizza-dog)")

	outputDirLabel := widget.NewLabel("Output directory: (current)")
	outputDirLabel.Wrapping = fyne.TextWrapWord

	selectDirBtn := widget.NewButton("Select Output Directory", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if uri != nil {
				outputDirLabel.SetText("Output directory: " + uri.Path())
			}
		}, state.Window)
	})

	// Options
	overwriteCheck := widget.NewCheck("Overwrite existing files", nil)
	
	// Action buttons
	receiveBtn := widget.NewButton("Receive", func() {
		// Placeholder for receive logic
	})
	receiveBtn.Importance = widget.HighImportance

	cancelBtn := widget.NewButton("Cancel", func() {
		// Placeholder for cancel logic
	})

	// Progress UI
	progressBar := widget.NewProgressBar()
	progressLabel := widget.NewLabel("Ready")
	fileListLabel := widget.NewLabel("")

	// Layout assembly
	codeBox := container.NewVBox(
		widget.NewLabel("Secret Code:"),
		codeEntry,
	)

	dirBox := container.NewVBox(
		selectDirBtn,
		outputDirLabel,
	)

	optionsBox := container.NewVBox(
		widget.NewLabel("Options:"),
		overwriteCheck,
	)

	actionBox := container.NewHBox(receiveBtn, cancelBtn)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Receive Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		codeBox,
		widget.NewSeparator(),
		dirBox,
		widget.NewSeparator(),
		optionsBox,
		widget.NewSeparator(),
		actionBox,
		widget.NewSeparator(),
		progressLabel,
		progressBar,
		fileListLabel,
	)

	return container.NewPadded(content)
}
