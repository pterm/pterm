# interactive_multiselect/demo

![Animation](https://vhs.charm.sh/vhs-6q1Hopg8iHBDoZb9e6yho4.gif)

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	var options []string
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// A few extra options that are easy to find with the fuzzy search filter.
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// The multiselect scrolls through long lists and supports fuzzy filtering:
	// just start typing to narrow down the options. By default, enter toggles
	// an option and tab confirms the selection.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```
