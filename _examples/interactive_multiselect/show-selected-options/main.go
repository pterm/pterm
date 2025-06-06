package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options.
	var options []string

	// Populate the options slice with 100 options.
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Use PTerm's interactive multiselect to present the options to the user and capture their selections.
	// The Show() method displays the options and waits for user input.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithShowSelectedOptions(true).
		Show()

	// Print the selected options, highlighted in green.
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
