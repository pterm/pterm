# coloring/override-default-printers

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a default error message with PTerm's built-in Error style.
	pterm.Error.Println("This is the default Error")

	// Override the default error prefix with a new text and style.
	pterm.Error.Prefix = pterm.Prefix{Text: "OVERRIDE", Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed)}

	// Print the error message again, this time with the overridden prefix.
	pterm.Error.Println("This is the default Error after the prefix was overridden")
}

```
