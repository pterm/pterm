# coloring/print-color-rgb

![Animation](https://vhs.charm.sh/vhs-4Iok4H7S6qYWGRtf5H2dTf.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// NewRGB creates a TrueColor foreground color; every RGB value can be
	// used as a printer directly.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")

	// Passing true as the optional last argument makes the color apply to
	// the background instead of the text.
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}
```
