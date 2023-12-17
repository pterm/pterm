package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define RGB colors for foreground and background.
	foregroundRGB := pterm.RGB{R: 187, G: 80, B: 0}
	backgroundRGB := pterm.RGB{R: 0, G: 50, B: 123}

	// Create a new RGB style with the defined foreground and background colors.
	rgbStyle := pterm.NewRGBStyle(foregroundRGB, backgroundRGB)

	// Print a string with the custom RGB style.
	rgbStyle.Println("This text is not styled.")

	// Add the 'Bold' option to the RGB style and print a string with this style.
	rgbStyle.AddOptions(pterm.Bold).Println("This text is bold.")

	// Add the 'Italic' option to the RGB style and print a string with this style.
	rgbStyle.AddOptions(pterm.Italic).Println("This text is italic.")
}
