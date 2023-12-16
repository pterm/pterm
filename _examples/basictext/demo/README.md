# basictext/demo

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The DefaultBasicText is a basic text printer provided by PTerm.
	// It is used to print text without any special formatting.
	pterm.DefaultBasicText.Println("Default basic text printer.")

	// The DefaultBasicText can be used in any context that requires a TextPrinter.
	// Here, we're using it with the LightMagenta function to color a portion of the text.
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")

	// The DefaultBasicText is also useful for resolving progress bars and spinners.
}

```
