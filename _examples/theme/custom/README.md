# theme/custom

![Animation](animation.gif)

```go
package main

import (
	"atomicgo.dev/color"

	"github.com/pterm/pterm"
)

func main() {
	// Start from a preset and change the slots you care about. Every printer
	// picks up the new theme, there is nothing else to restyle.
	theme := pterm.ThemeMacchiato
	theme.Primary = color.RGB(0xf5, 0xa9, 0x7f) // peach
	theme.Success = color.RGB(0x8b, 0xd5, 0xca) // teal

	pterm.SetTheme(theme)

	pterm.Success.Println("Deployed to production")
	pterm.Callout.Tip("Theme").Println("Primary is now peach and success is teal.")

	// Tones can also be fixed to a color that ignores the theme.
	pterm.Println(pterm.Fixed(color.RGB(0xc6, 0xa0, 0xf6)).Sprint("Always mauve"))
}
```
