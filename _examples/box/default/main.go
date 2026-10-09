package main

import "github.com/pterm/pterm"

func main() {
	// The box sizes itself to fit whatever it prints.
	pterm.DefaultBox.Println("Hello, World!")
}
