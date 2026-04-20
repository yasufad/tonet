package ui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/yasufad/tonet/src/croc"
	"github.com/yasufad/tonet/src/models"
)

func makeReceiveTab(state *AppState) fyne.CanvasObject {
	outputDir, _ := os.Getwd()

	codeEntry := widget.NewEntry()
	codeEntry.SetPlaceHolder("Enter receive code (e.g., fast-pizza-dog)")

	outputDirLabel := widget.NewLabel("Output: " + outputDir)
	outputDirLabel.Wrapping = fyne.TextWrapWord

	selectDirBtn := widget.NewButton("Select Output Directory", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if uri != nil {
				outputDir = uri.Path()
				outputDirLabel.SetText("Output: " + outputDir)
			}
		}, state.Window)
	})

	overwriteCheck := widget.NewCheck("Overwrite existing files", nil)
	stdoutCheck := widget.NewCheck("Redirect to stdout", nil)
	onlyLocalCheck := widget.NewCheck("Force local connections only", nil)

	progressBar := widget.NewProgressBar()
	progressLabel := widget.NewLabel("Ready")

	var receiveBtn *widget.Button
	receiveBtn = widget.NewButton("Receive", func() {
		code := codeEntry.Text
		if code == "" {
			dialog.ShowError(fmt.Errorf("please enter a receive code"), state.Window)
			return
		}

		receiveBtn.Disable()
		progressLabel.SetText("Connecting...")

		go func() {
			currentDir, _ := os.Getwd()
			os.Chdir(outputDir)
			defer os.Chdir(currentDir)

			settings := loadSettings()

			opts := croc.Options{
				IsSender:      false,
				SharedSecret:  code,
				RelayAddress:  settings.RelayAddress,
				RelayAddress6: settings.RelayAddress6,
				RelayPassword: settings.RelayPassword,
				Curve:         settings.Curve,
				Overwrite:     overwriteCheck.Checked,
				Stdout:        stdoutCheck.Checked,
				OnlyLocal:     onlyLocalCheck.Checked,
				NoPrompt:      true,
			}

			if opts.RelayAddress == "" {
				opts.RelayAddress = models.DEFAULT_RELAY
			}
			if opts.RelayAddress6 == "" {
				opts.RelayAddress6 = models.DEFAULT_RELAY6
			}
			if opts.RelayPassword == "" {
				opts.RelayPassword = models.DEFAULT_PASSPHRASE
			}
			if opts.Curve == "" {
				opts.Curve = "p256"
			}

			cr, err := croc.New(opts)
			if err != nil {
				dialog.ShowError(err, state.Window)
				receiveBtn.Enable()
				return
			}

			progressLabel.SetText("Receiving...")
			err = cr.Receive()
			if err != nil {
				dialog.ShowError(err, state.Window)
			} else {
				progressLabel.SetText("Transfer complete!")
				dialog.ShowInformation("Success", "Files received successfully!", state.Window)
			}
			receiveBtn.Enable()
		}()
	})
	receiveBtn.Importance = widget.HighImportance

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
		stdoutCheck,
		onlyLocalCheck,
	)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Receive Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		codeBox,
		widget.NewSeparator(),
		dirBox,
		widget.NewSeparator(),
		optionsBox,
		widget.NewSeparator(),
		receiveBtn,
		widget.NewSeparator(),
		progressLabel,
		progressBar,
	)

	return container.NewPadded(content)
}
