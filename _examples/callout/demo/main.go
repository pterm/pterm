package main

import "github.com/pterm/pterm"

func main() {
	pterm.Callout.Note("Note").Println("Neutral information that supports the command.")
	pterm.Println()
	pterm.Callout.Warn("Heads up").Println("Something to double check before continuing.")
	pterm.Println()
	pterm.Callout.Tip("Done").Println("This is the good outcome.")
	pterm.Println()
	pterm.Callout.Danger("Error").Println("Something went wrong while reading the config.")
}
