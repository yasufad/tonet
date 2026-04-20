package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/yasufad/tonet/src/croc"
	"github.com/yasufad/tonet/src/models"
	"github.com/yasufad/tonet/src/utils"
)

func makeSendTab(state *AppState) fyne.CanvasObject {
	var selectedFiles []string

	selectedFilesLabel := widget.NewLabel("No files selected")
	selectedFilesLabel.Wrapping = fyne.TextWrapWord

	selectFilesBtn := widget.NewButton("Select Files", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if reader != nil {
				selectedFiles = []string{reader.URI().Path()}
				selectedFilesLabel.SetText("Selected: " + reader.URI().Name())
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
				selectedFiles = []string{uri.Path()}
				selectedFilesLabel.SetText("Selected: " + uri.Name())
			}
		}, state.Window)
	})

	hashAlgoSelect := widget.NewSelect([]string{"xxhash", "imohash", "md5"}, nil)
	hashAlgoSelect.SetSelected("xxhash")

	zipFolderCheck := widget.NewCheck("Zip folder", nil)
	zipFolderCheck.SetChecked(true)

	gitIgnoreCheck := widget.NewCheck("Respect .gitignore", nil)

	customCodeCheck := widget.NewCheck("Use custom code", nil)
	customCodeEntry := widget.NewEntry()
	customCodeEntry.SetPlaceHolder("Leave empty for auto-generated code")
	customCodeEntry.Disable()

	customCodeCheck.OnChanged = func(checked bool) {
		if checked {
			customCodeEntry.Enable()
		} else {
			customCodeEntry.Disable()
		}
	}

	// Advanced options
	noCompressCheck := widget.NewCheck("Disable compression", nil)
	disableLocalCheck := widget.NewCheck("Disable local relay", nil)
	noMultiCheck := widget.NewCheck("Disable multiplexing", nil)
	onlyLocalCheck := widget.NewCheck("Force local connections only", nil)
	showQrCheck := widget.NewCheck("Show QR code for mobile receive", nil)
	throttleEntry := widget.NewEntry()
	throttleEntry.SetPlaceHolder("Upload speed limit (e.g. 500k, 1m)")
	excludeEntry := widget.NewEntry()
	excludeEntry.SetPlaceHolder("Exclude patterns (comma-separated, e.g. node_modules,.git)")

	codeLabel := widget.NewLabel("")
	codeLabel.Wrapping = fyne.TextWrapWord

	progressBar := widget.NewProgressBar()
	progressLabel := widget.NewLabel("Ready")

	var sendBtn *widget.Button
	sendBtn = widget.NewButton("Send", func() {
		if len(selectedFiles) == 0 {
			dialog.ShowError(fmt.Errorf("no files selected"), state.Window)
			return
		}

		sendBtn.Disable()
		progressLabel.SetText("Preparing files...")

		go func() {
			settings := loadSettings()

			customCode := customCodeEntry.Text
			if customCode == "" || !customCodeCheck.Checked {
				customCode = utils.GetRandomName()
			}

			opts := croc.Options{
				IsSender:       true,
				RelayAddress:   settings.RelayAddress,
				RelayAddress6:  settings.RelayAddress6,
				RelayPorts:     []string{"9009", "9010", "9011", "9012", "9013"},
				RelayPassword:  settings.RelayPassword,
				SharedSecret:   customCode,
				HashAlgorithm:  hashAlgoSelect.Selected,
				ZipFolder:      zipFolderCheck.Checked,
				GitIgnore:      gitIgnoreCheck.Checked,
				Curve:          settings.Curve,
				NoCompress:     noCompressCheck.Checked,
				DisableLocal:   disableLocalCheck.Checked,
				NoMultiplexing: noMultiCheck.Checked,
				OnlyLocal:      onlyLocalCheck.Checked,
				ShowQrCode:     showQrCheck.Checked,
				ThrottleUpload: throttleEntry.Text,
				NoPrompt:       true,
			}

			// Parse exclude patterns
			if excludeEntry.Text != "" {
				for _, v := range strings.Split(excludeEntry.Text, ",") {
					v = strings.TrimSpace(v)
					if v != "" {
						opts.Exclude = append(opts.Exclude, v)
					}
				}
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

			fileInfos, emptyFolders, totalFolders, err := croc.GetFilesInfo(selectedFiles, opts.ZipFolder, opts.GitIgnore, []string{})
			if err != nil {
				dialog.ShowError(err, state.Window)
				sendBtn.Enable()
				return
			}

			codeLabel.SetText("Code: " + opts.SharedSecret)
			progressLabel.SetText("Connecting to relay...")

			cr, err := croc.New(opts)
			if err != nil {
				dialog.ShowError(err, state.Window)
				sendBtn.Enable()
				return
			}

			progressLabel.SetText("Sending...")
			err = cr.Send(fileInfos, emptyFolders, totalFolders)
			if err != nil {
				dialog.ShowError(err, state.Window)
			} else {
				progressLabel.SetText("Transfer complete!")
				dialog.ShowInformation("Success", "Files sent successfully!", state.Window)
			}
			sendBtn.Enable()
		}()
	})
	sendBtn.Importance = widget.HighImportance

	filesBox := container.NewHBox(selectFilesBtn, selectFolderBtn)
	optionsBox := container.NewVBox(
		widget.NewLabel("Options:"),
		container.NewHBox(widget.NewLabel("Hash:"), hashAlgoSelect),
		zipFolderCheck,
		gitIgnoreCheck,
		widget.NewSeparator(),
		customCodeCheck,
		customCodeEntry,
		widget.NewSeparator(),
		widget.NewLabel("Advanced:"),
		noCompressCheck,
		disableLocalCheck,
		noMultiCheck,
		onlyLocalCheck,
		showQrCheck,
		widget.NewLabel("Upload speed limit:"),
		throttleEntry,
		widget.NewLabel("Exclude:"),
		excludeEntry,
	)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Send Files", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		filesBox,
		selectedFilesLabel,
		widget.NewSeparator(),
		optionsBox,
		widget.NewSeparator(),
		sendBtn,
		widget.NewSeparator(),
		codeLabel,
		progressLabel,
		progressBar,
	)

	return container.NewPadded(content)
}
