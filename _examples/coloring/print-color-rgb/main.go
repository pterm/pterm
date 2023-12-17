package main

import "github.com/pterm/pterm"

func main() {
	// Create a new RGB color with values 178, 44, 199.
	// This color will be used for the text.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")

	// Create a new RGB color with values 15, 199, 209.
	// This color will be used for the text.
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")

	// Create a new RGB color with values 201, 144, 30.
	// This color will be used for the background.
	// The 'true' argument indicates that the color is for the background.
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}
