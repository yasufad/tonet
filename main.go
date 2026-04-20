package main

//go:generate go run src/install/updateversion.go
//go:generate git commit -am "bump $VERSION"
//go:generate git tag -af v$VERSION -m "v$VERSION"

import (
	"fmt"
	"os"

	"github.com/yasufad/tonet/src/cli"
	"github.com/yasufad/tonet/src/ui"
)

func main() {
	if len(os.Args) > 1 {
		// Run original CLI
		if err := cli.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// Boot Fyne GUI
		ui.Run()
	}
}
