# coloring/print-color-rgb-style

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	foregroundRGB := pterm.RGB{R: 187, G: 80, B: 0}
	backgroundRGB := pterm.RGB{R: 0, G: 50, B: 123}

	// Print string with a custom foreground and background RGB color.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).Println("This text is not styled.")

	// Print string with a custom foreground and background RGB color and style bold.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).AddOptions(pterm.Bold).Println("This text is bold.")

	// Print string with a custom foreground and background RGB color and style italic.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).AddOptions(pterm.Italic).Println("This text is italic.")
}

```
