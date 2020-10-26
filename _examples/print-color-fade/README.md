# print-color-fade

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("Please note that RGB colors only work in Terminals which support TrueColor.")

	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)
	for i := 0; i < pterm.GetTerminalHeight()-1; i++ {
		from.Fade(0, float32(pterm.GetTerminalHeight()), float32(i), to).Println("Hello, World!")
	}
}

```
