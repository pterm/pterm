package main

import "github.com/pterm/pterm"

func main() {
	// A message with more than one line turns into a callout. The first
	// line becomes the title.
	pterm.Error.Println("Unknown command disk\nRun `mjw --help` to see available commands.")

	pterm.Println()

	// WithTitle sets the title explicitly and keeps the whole message as
	// the body.
	pterm.Warning.WithTitle("Heads up").Println("Nothing here can be undone.")
}
