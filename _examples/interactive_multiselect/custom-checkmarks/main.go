package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Populate the options slice with 5 options
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Create a new interactive multiselect printer with the options
	// Disable the filter and define the checkmark symbols
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithCheckmark(&pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")})

	// Show the interactive multiselect and get the selected options
	selectedOptions, _ := printer.Show()

	// Print the selected options
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
