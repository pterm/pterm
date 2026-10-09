package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The continue prompt offers "yes", "no", "all" and "cancel" by default.
	// Each option is chosen by typing its first letter; pressing enter picks
	// the default answer (the first option, unless changed).
	result, _ := pterm.DefaultInteractiveContinue.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
