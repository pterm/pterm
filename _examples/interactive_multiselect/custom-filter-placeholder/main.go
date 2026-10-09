package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	var options []string
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// A few extra options that are easy to find with the fuzzy search filter.
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// WithFilterInputPlaceholder replaces the hint text that is shown in the
	// filter input while it is still empty.
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	selectedOptions, _ := printer.Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
