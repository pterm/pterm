package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The default value is shown as a pre-filled suggestion. Pressing enter
	// right away returns it; typing anything replaces it.
	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("Some default value").Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
