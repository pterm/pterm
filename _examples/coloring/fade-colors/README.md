# coloring/fade-colors

![Animation](https://vhs.charm.sh/vhs-6uuzyL2JxFTfYSemtD7OiA.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// RGB colors need a TrueColor terminal; on anything less PTerm
	// downsamples them to the closest supported color.
	pterm.Info.Println("RGB colors only work in Terminals which support TrueColor.")

	startColor := pterm.NewRGB(0, 255, 255) // cyan
	endColor := pterm.NewRGB(255, 0, 255)   // magenta

	// Spread the gradient over the visible terminal height, one line per step.
	terminalHeight := pterm.GetTerminalHeight()

	for i := 0; i < terminalHeight-2; i++ {
		// Fade interpolates between the two colors; the factor 0..1 selects
		// the position on the gradient.
		fadeFactor := float32(i) / float32(terminalHeight-2)
		currentColor := startColor.Fade(0, 1, fadeFactor, endColor)

		currentColor.Println("Hello, World!")
	}
}
```
