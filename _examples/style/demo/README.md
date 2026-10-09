# style/demo

![Animation](https://vhs.charm.sh/vhs-5fzvhSDBpEvbWqcS5BrJhl.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A Style combines any number of colors and text options.
	// It can be reused anywhere PTerm accepts a style.
	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)
	secondary := pterm.NewStyle(pterm.FgLightGreen, pterm.BgWhite, pterm.Italic)

	// Styles are also TextPrinters, so they can print directly.
	primary.Println("Hello, World!")
	secondary.Println("Hello, World!")
}
```
