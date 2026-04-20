package ui

import (
	"encoding/json"
	"os"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/yasufad/tonet/src/models"
	"github.com/yasufad/tonet/src/utils"
)

type Settings struct {
	RelayAddress  string `json:"relay_address"`
	RelayAddress6 string `json:"relay_address6"`
	RelayPassword string `json:"relay_password"`
	Curve         string `json:"curve"`
	Socks5Proxy   string `json:"socks5_proxy"`
	HttpProxy     string `json:"http_proxy"`
}

func makeSettingsTab(state *AppState) fyne.CanvasObject {
	settings := loadSettings()

	relayEntry := widget.NewEntry()
	relayEntry.SetText(settings.RelayAddress)
	relayEntry.SetPlaceHolder(models.DEFAULT_RELAY)

	relay6Entry := widget.NewEntry()
	relay6Entry.SetText(settings.RelayAddress6)
	relay6Entry.SetPlaceHolder(models.DEFAULT_RELAY6)

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetText(settings.RelayPassword)
	passwordEntry.SetPlaceHolder(models.DEFAULT_PASSPHRASE)

	curveSelect := widget.NewSelect([]string{"p256", "p384", "p521", "siec"}, nil)
	if settings.Curve != "" {
		curveSelect.SetSelected(settings.Curve)
	} else {
		curveSelect.SetSelected("p256")
	}

	socks5Entry := widget.NewEntry()
	socks5Entry.SetText(settings.Socks5Proxy)
	socks5Entry.SetPlaceHolder("e.g., 127.0.0.1:9050")

	httpProxyEntry := widget.NewEntry()
	httpProxyEntry.SetText(settings.HttpProxy)
	httpProxyEntry.SetPlaceHolder("e.g., http://proxy:8080")

	saveBtn := widget.NewButton("Save Settings", func() {
		newSettings := Settings{
			RelayAddress:  relayEntry.Text,
			RelayAddress6: relay6Entry.Text,
			RelayPassword: passwordEntry.Text,
			Curve:         curveSelect.Selected,
			Socks5Proxy:   socks5Entry.Text,
			HttpProxy:     httpProxyEntry.Text,
		}

		if err := saveSettings(newSettings); err != nil {
			dialog.ShowError(err, state.Window)
		} else {
			dialog.ShowInformation("Success", "Settings saved successfully!", state.Window)
		}
	})
	saveBtn.Importance = widget.HighImportance

	resetBtn := widget.NewButton("Reset to Defaults", func() {
		relayEntry.SetText(models.DEFAULT_RELAY)
		relay6Entry.SetText(models.DEFAULT_RELAY6)
		passwordEntry.SetText(models.DEFAULT_PASSPHRASE)
		curveSelect.SetSelected("p256")
		socks5Entry.SetText("")
		httpProxyEntry.SetText("")
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle("Settings", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Relay Configuration:"),
		widget.NewLabel("IPv4 Relay Address:"),
		relayEntry,
		widget.NewLabel("IPv6 Relay Address:"),
		relay6Entry,
		widget.NewLabel("Relay Password:"),
		passwordEntry,
		widget.NewSeparator(),
		widget.NewLabel("Encryption Curve:"),
		curveSelect,
		widget.NewSeparator(),
		widget.NewLabel("Proxy Settings:"),
		widget.NewLabel("SOCKS5 Proxy:"),
		socks5Entry,
		widget.NewLabel("HTTP Proxy:"),
		httpProxyEntry,
		widget.NewSeparator(),
		container.NewHBox(saveBtn, resetBtn),
	)

	return container.NewPadded(container.NewVScroll(content))
}

func getSettingsFile() string {
	configDir, err := utils.GetConfigDir(true)
	if err != nil {
		return ""
	}
	return path.Join(configDir, "tonet-settings.json")
}

func loadSettings() Settings {
	settings := Settings{
		RelayAddress:  models.DEFAULT_RELAY,
		RelayAddress6: models.DEFAULT_RELAY6,
		RelayPassword: models.DEFAULT_PASSPHRASE,
		Curve:         "p256",
	}

	file := getSettingsFile()
	if file == "" {
		return settings
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return settings
	}

	json.Unmarshal(data, &settings)
	return settings
}

func saveSettings(settings Settings) error {
	file := getSettingsFile()
	if file == "" {
		return nil
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
