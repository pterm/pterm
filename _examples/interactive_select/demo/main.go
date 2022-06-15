package main

import (
	"github.com/pterm/pterm"
)

func main() {
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Option 1", "Option 2", "Option 3", "Option 4"}).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
