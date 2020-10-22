# print-basic-letters

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBasicText.Println("Default basic letters printer.")
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
	pterm.DefaultBasicText.Println("For example to resolve progressbars and spinners.")
}

```
