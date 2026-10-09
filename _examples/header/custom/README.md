# header/custom

![Animation](https://vhs.charm.sh/vhs-5c8QzHYrjymhxB5DTG6XEF.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Restyle the default header on the fly. The margin adds horizontal
	// padding on both sides of the text.
	pterm.DefaultHeader.WithMargin(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("This is a custom header!")

	// Alternatively, build a HeaderPrinter from scratch instead of deriving
	// from DefaultHeader.
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	newHeader.Println("This is a custom header!")
}
```
