package main

import (
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	// Define RGB values for gradient points.
	startColor := pterm.NewRGB(0, 255, 255)
	firstPoint := pterm.NewRGB(255, 0, 255)
	secondPoint := pterm.NewRGB(255, 0, 0)
	thirdPoint := pterm.NewRGB(0, 255, 0)
	endColor := pterm.NewRGB(255, 255, 255)

	// Define the string to be printed.
	str := "RGB colors only work in Terminals which support TrueColor."
	strs := strings.Split(str, "")

	// Initialize an empty string for the faded info.
	var fadeInfo string

	// Loop over the string length to create a gradient effect.
	for i := 0; i < len(str); i++ {
		// Append each character of the string with a faded color to the info string.
		fadeInfo += startColor.Fade(0, float32(len(str)), float32(i), firstPoint).Sprint(strs[i])
	}

	// Print the info string with gradient effect.
	pterm.Info.Println(fadeInfo)

	// Get the terminal height.
	terminalHeight := pterm.GetTerminalHeight()

	// Loop over the terminal height to print "Hello, World!" with a gradient effect.
	for i := 0; i < terminalHeight-2; i++ {
		// Print the string with a color that fades from startColor to endColor.
		startColor.Fade(0, float32(terminalHeight-2), float32(i), firstPoint, secondPoint, thirdPoint, endColor).Println("Hello, World!")
	}
}
