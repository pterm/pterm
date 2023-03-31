# coloring/print-color-rgb

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print strings with a custom RGB color.
	// NOTICE: This only works with terminals which support TrueColor.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}

```
