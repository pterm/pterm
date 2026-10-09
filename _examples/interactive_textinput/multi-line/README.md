# interactive_textinput/multi-line

![Animation](https://vhs.charm.sh/vhs-6BcAPqMvP68DHVX0tCfZg5.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// In multi-line mode, enter inserts a new line and tab submits the input.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	result, _ := textInput.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
```
