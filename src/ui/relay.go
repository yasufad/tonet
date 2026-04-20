package ui

import (
	"context"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yasufad/tonet/src/models"
	"github.com/yasufad/tonet/src/tcp"
)

func makeRelayTab(state *AppState) fyne.CanvasObject {
	var cancel context.CancelFunc

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Leave empty for all interfaces")

	portEntry := widget.NewEntry()
	portEntry.SetText("9009")
	portEntry.SetPlaceHolder("9009")

	transfersEntry := widget.NewEntry()
	transfersEntry.SetText("4")
	transfersEntry.SetPlaceHolder("4")

	statusLabel := widget.NewLabel("Relay stopped")
	logLabel := widget.NewLabel("")
	logLabel.Wrapping = fyne.TextWrapWord

	var toggleBtn *widget.Button
	toggleBtn = widget.NewButton("Start Relay", func() {
		if cancel != nil {
			// Stop relay
			cancel()
			cancel = nil
			statusLabel.SetText("Relay stopped")
			toggleBtn.SetText("Start Relay")
			toggleBtn.Importance = widget.MediumImportance
			return
		}

		// Start relay
		basePort, err := strconv.Atoi(portEntry.Text)
		if err != nil || basePort < 1 || basePort > 65530 {
			statusLabel.SetText("Error: Invalid port number")
			return
		}

		transfers, err := strconv.Atoi(transfersEntry.Text)
		if err != nil || transfers < 2 {
			statusLabel.SetText("Error: Transfers must be at least 2")
			return
		}

		host := hostEntry.Text
		ports := make([]string, transfers)
		for i := 0; i < transfers; i++ {
			ports[i] = strconv.Itoa(basePort + i)
		}

		ctx, cancelFunc := context.WithCancel(context.Background())
		cancel = cancelFunc

		statusLabel.SetText(fmt.Sprintf("Relay running on ports %s-%s", ports[0], ports[len(ports)-1]))
		toggleBtn.SetText("Stop Relay")
		toggleBtn.Importance = widget.HighImportance

		go func() {
			// Start transfer ports
			for i := 1; i < len(ports); i++ {
				port := ports[i]
				go func(p string) {
					tcp.Run("info", host, p, models.DEFAULT_PASSPHRASE)
				}(port)
			}

			// Start main port (blocking)
			tcpPorts := ""
			for i := 1; i < len(ports); i++ {
				if i > 1 {
					tcpPorts += ","
				}
				tcpPorts += ports[i]
			}
			tcp.Run("info", host, ports[0], models.DEFAULT_PASSPHRASE, tcpPorts)

			<-ctx.Done()
		}()
	})
	toggleBtn.Importance = widget.MediumImportance

	hostBox := container.NewVBox(
		widget.NewLabel("Host (optional):"),
		hostEntry,
	)

	portBox := container.NewVBox(
		widget.NewLabel("Base Port:"),
		portEntry,
	)

	transfersBox := container.NewVBox(
		widget.NewLabel("Number of Transfer Ports:"),
		transfersEntry,
	)

	content := container.NewVBox(
		widget.NewLabelWithStyle("Relay Server", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		hostBox,
		portBox,
		transfersBox,
		widget.NewSeparator(),
		toggleBtn,
		widget.NewSeparator(),
		statusLabel,
		logLabel,
	)

	return container.NewPadded(content)
}
