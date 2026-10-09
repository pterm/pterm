# interactive_multiselect/custom-checkmarks

![Animation](https://vhs.charm.sh/vhs-1AgHtATkrEGGCyLepQ380I.gif)

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	var options []string
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// The checkmark in front of each option can be any string, here a green
	// plus for selected and a red minus for unselected options. The fuzzy
	// search filter is disabled to keep the list short and static.
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithCheckmark(&pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")})

	selectedOptions, _ := printer.Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```
