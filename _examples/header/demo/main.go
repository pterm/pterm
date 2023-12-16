package main

import "github.com/pterm/pterm"

func main() {
	// Print a default header.
	// This uses the default settings of PTerm to print a header.
	pterm.DefaultHeader.Println("This is the default header!")

	// Print a spacer line for better readability.
	pterm.Println()

	// Print a full-width header.
	// This uses the WithFullWidth() option of PTerm to print a header that spans the full width of the terminal.
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}
