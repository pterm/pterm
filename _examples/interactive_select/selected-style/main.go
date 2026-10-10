package main

import "github.com/pterm/pterm"

func main() {
	options := []string{"Go", "Rust", "TypeScript", "Python", "Zig"}

	// The highlighted option gets its own style, so it stands out from the
	// rest of the list. Without it, every option uses OptionStyle.
	selectedOption, _ := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithSelectedOptionStyle(pterm.NewStyle(pterm.FgLightCyan, pterm.Bold)).
		Show("Pick a language")

	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
