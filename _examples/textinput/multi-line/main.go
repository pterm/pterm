package main

import "github.com/pterm/pterm"

func main() {
	// In multi-line mode enter starts a new line and ctrl+d submits. Long
	// lines wrap at word boundaries, and pasted text keeps its line breaks.
	message, _ := pterm.TextInput.
		WithMultiLine().
		Show("Commit message")

	pterm.Info.Printfln("Committing %d characters", len(message))
}
