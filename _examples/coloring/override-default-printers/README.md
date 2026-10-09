# coloring/override-default-printers

![Animation](https://vhs.charm.sh/vhs-5uqfqe2Y4nIt6DAahZxgrR.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Error.Println("This is the default Error")

	// The default printers are package-level variables, so their fields can
	// be changed directly. This swaps the Error prefix text and style for
	// every subsequent pterm.Error call in the program.
	pterm.Error.Prefix = pterm.Prefix{Text: "OVERRIDE", Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed)}

	pterm.Error.Println("This is the default Error after the prefix was overridden")
}
```
