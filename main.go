package main

//go:generate go run src/install/updateversion.go
//go:generate git commit -am "bump $VERSION"
//go:generate git tag -af v$VERSION -m "v$VERSION"

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/schollz/croc/v10/src/cli"
	"github.com/schollz/croc/v10/src/ui"
	"github.com/schollz/croc/v10/src/utils"
)

func main() {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Determine whether to run CLI or GUI
	if len(os.Args) > 1 {
		// Run original CLI
		go func() {
			if err := cli.Run(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			utils.RemoveMarkedFiles()
			os.Exit(0)
		}()
	} else {
		// Boot Fyne GUI
		go func() {
			ui.Run()
			utils.RemoveMarkedFiles()
			os.Exit(0)
		}()
	}

	// Wait for a termination signal
	<-sigs
	utils.RemoveMarkedFiles()
	os.Exit(0)
}
