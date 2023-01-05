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

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	printer.Checkmark = &pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")}
	selectedOptions, _ := printer.Show()
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
