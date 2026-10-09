package main

import (
	"fmt"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

func main() {
	var options []string
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// By default, enter toggles an option and tab confirms the selection.
	// Both keys can be rebound, here to space for toggling and enter for
	// confirming. The filter is disabled so typing is not needed.
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space)

	selectedOptions, _ := printer.Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
