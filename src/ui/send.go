package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func makeSendTab(state *AppState) fyne.CanvasObject {
	// UI Elements
	selectedFilesLabel := widget.NewLabel("No files selected")
	selectedFilesLabel.Wrapping = fyne.TextWrapWord

	// File selection buttons
	selectFilesBtn := widget.NewButton("Select Files", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if reader != nil {
				selectedFilesLabel.SetText("Selected file: " + reader.URI().Name())
			}
		}, state.Window)
	})

	selectFolderBtn := widget.NewButton("Select Folder", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if uri != nil {
				selectedFilesLabel.SetText("Selected folder: " + uri.Name())
			}
		}, state.Window)
	})

	// Options
	hashAlgoSelect := widget.NewSelect([]string{"xxhash", "imohash", "md5"}, nil)
	hashAlgoSelect.SetSelected("xxhash")
	
	zipFolderCheck := widget.NewCheck("Zip folder", nil)
	zipFolderCheck.SetChecked(true)

	// Action buttons
	sendBtn := widget.NewButton("Send", func() {
		// Placeholder for sending logic
	})
	sendBtn.Importance = widget.HighImportance

	cancelBtn := widget.NewButton("Cancel", func() {
		// Placeholder for cancel logic
	})

	// Progress UI
	progressBar := widget.NewProgressBar()
	progressLabel := widget.NewLabel("Ready")

	// Layout assembly
	filesBox := container.NewHBox(selectFilesBtn, selectFolderBtn)
	optionsBox := container.NewVBox(
		widget.NewLabel("Options:"),
		container.NewHBox(widget.NewLabel("Hash Algorithm:"), hashAlgoSelect),
		zipFolderCheck,
	)
	actionBox := container.NewHBox(sendBtn, cancelBtn)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Send Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		filesBox,
		selectedFilesLabel,
		widget.NewSeparator(),
		optionsBox,
		widget.NewSeparator(),
		actionBox,
		widget.NewSeparator(),
		progressLabel,
		progressBar,
	)

	return container.NewPadded(content)
}
