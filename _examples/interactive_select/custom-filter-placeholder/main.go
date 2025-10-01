package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Generate 100 options and add them to the options slice
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Generate 5 additional options with a specific message and add them to the options slice
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Create a new interactive select printer with custom filter placeholder text
	// The WithFilterInputPlaceholder method allows you to customize the placeholder text shown when filtering
	printer := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := printer.Show()

	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
