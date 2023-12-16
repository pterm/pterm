# coloring/fade-colors

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Print an informational message.
	pterm.Info.Println("RGB colors only work in Terminals which support TrueColor.")

	// Define the start and end points for the color gradient.
	startColor := pterm.NewRGB(0, 255, 255) // Cyan
	endColor := pterm.NewRGB(255, 0, 255)   // Magenta

	// Get the terminal height to determine the gradient range.
	terminalHeight := pterm.GetTerminalHeight()

	// Loop over the range of the terminal height to create a color gradient.
	for i := 0; i < terminalHeight-2; i++ {
		// Calculate the fade factor for the current step in the gradient.
		fadeFactor := float32(i) / float32(terminalHeight-2)

		// Create a color that represents the current step in the gradient.
		currentColor := startColor.Fade(0, 1, fadeFactor, endColor)

		// Print a string with the current color.
		currentColor.Println("Hello, World!")
	}
}

```
