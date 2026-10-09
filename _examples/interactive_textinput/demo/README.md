# interactive_textinput/demo

![Animation](https://vhs.charm.sh/vhs-7xu04W1sjODqqV4xW1bQrY.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The text input is single-line by default; enter submits the input.
	result, _ := pterm.DefaultInteractiveTextInput.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
```
