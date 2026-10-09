### interactive_select/demo

![Animation](https://vhs.charm.sh/vhs-1ucPzCtHk2GpUpPvRTb5wp.gif)

<details>

<summary>SHOW SOURCE</summary>

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

	// The select prompt scrolls through long lists and supports fuzzy
	// filtering: just start typing to narrow down the options. Enter picks
	// the highlighted option.
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
```

</details>

### interactive_select/custom-filter-placeholder

![Animation](https://vhs.charm.sh/vhs-4KMwNGwkyEpA4mqOGnZY2I.gif)

<details>

<summary>SHOW SOURCE</summary>

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

	// WithFilterInputPlaceholder replaces the hint text that is shown in the
	// filter input while it is still empty.
	printer := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	selectedOption, _ := printer.Show()

	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
```

</details>

