package main

import (
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

func main() {
	// Space toggles and enter confirms by default. Here enter toggles and
	// tab confirms instead, and the checkboxes are drawn in brackets. The
	// key hints at the bottom follow the keys.
	services, _ := pterm.Multiselect.
		WithOptions([]string{"api", "worker", "scheduler", "web"}).
		WithKeySelect(keys.Enter).
		WithKeyConfirm(keys.Tab).
		WithCheckmark("[x]", "[ ]").
		Show("Which services should restart?")

	pterm.Info.Printfln("Restarting %v", services)
}
