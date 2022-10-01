# interactive_multiselect/custom-keys

![Animation](animation.svg)

```go
package main

import (
	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func main() {
	var options []string

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	selectedOptions, _ := printer.Show()
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}

```
