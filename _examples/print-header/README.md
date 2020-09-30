# print-header

![Animation](animation.svg)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	tick := time.Tick(time.Second * 2)

	// Print with the default HeaderPrinter
	pterm.Header.Println("This is the default header style")

	<-tick // Wait

	// Print a custom header
	pterm.Header.SetFullWidth().SetTextStyle(pterm.FgDarkGray).SetBackgroundStyle(pterm.BgLightMagenta).Println("Hello, World!")

	<-tick // Wait

	// Create a custom HeaderPrinter
	customHeaderPrinter := pterm.HeaderPrinter{
		TextStyle:       pterm.Style{pterm.FgLightRed},
		BackgroundStyle: pterm.Style{pterm.BgGreen},
		Margin:          15,
	}
	// Use custom Header printer
	customHeaderPrinter.Println("This is a custom header.")
}

```
