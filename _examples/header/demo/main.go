package main

import "github.com/pterm/pterm"

func main() {
	// A header opens the output of a program: its name in a badge, what is
	// running and a short muted description, with a rule underneath.
	pterm.Header.
		WithBadge("MJW").
		WithSubtitle("Design system gallery").
		Println("debug ui")

	pterm.Text.Println("Everything below the header belongs to this run.")
}
