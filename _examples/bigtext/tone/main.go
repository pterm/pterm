package main

import "github.com/pterm/pterm"

func main() {
	// One tone instead of the gradient.
	pterm.BigText.WithTone(pterm.ToneSuccess).Println("Done")
}
