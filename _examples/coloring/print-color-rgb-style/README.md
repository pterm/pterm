# coloring/print-color-rgb-style

![Animation](https://vhs.charm.sh/vhs-78d44TVBSv37OO95sBQ22Z.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	foregroundRGB := pterm.RGB{R: 187, G: 80, B: 0}
	backgroundRGB := pterm.RGB{R: 0, G: 50, B: 123}

	// NewRGBStyle pairs a TrueColor foreground with a background color.
	rgbStyle := pterm.NewRGBStyle(foregroundRGB, backgroundRGB)

	rgbStyle.Println("This text is not styled.")

	// AddOptions returns a new style, so the bold and italic lines below are
	// independent of each other and of rgbStyle.
	rgbStyle.AddOptions(pterm.Bold).Println("This text is bold.")
	rgbStyle.AddOptions(pterm.Italic).Println("This text is italic.")
}
```
