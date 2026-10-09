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

	// WithShowSelectedOptions lists the current selection above the prompt,
	// which is handy when the chosen options are scrolled out of view.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithShowSelectedOptions(true).
		Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
