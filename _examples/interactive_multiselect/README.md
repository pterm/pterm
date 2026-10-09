### interactive_multiselect/demo

![Animation](https://vhs.charm.sh/vhs-6q1Hopg8iHBDoZb9e6yho4.gif)

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

	// The multiselect scrolls through long lists and supports fuzzy filtering:
	// just start typing to narrow down the options. By default, enter toggles
	// an option and tab confirms the selection.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/custom-checkmarks

![Animation](https://vhs.charm.sh/vhs-1AgHtATkrEGGCyLepQ380I.gif)

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

</details>

### interactive_multiselect/custom-filter-placeholder

![Animation](https://vhs.charm.sh/vhs-5T78j3vXVgR4xmDHlfSgMZ.gif)

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
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	selectedOptions, _ := printer.Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/custom-keys

![Animation](https://vhs.charm.sh/vhs-3X9IsnHV4BcVJfjuKiDbTp.gif)

<details>

<summary>SHOW SOURCE</summary>

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

	// By default, enter toggles an option and tab confirms the selection.
	// Both keys can be rebound, here to space for toggling and enter for
	// confirming. The filter is disabled so typing is not needed.
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space)

	selectedOptions, _ := printer.Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/show-selected-options

![Animation](https://vhs.charm.sh/vhs-4jmmHyWj8Zkw7WK7f2zf2D.gif)

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

	// WithShowSelectedOptions lists the current selection above the prompt,
	// which is handy when the chosen options are scrolled out of view.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithShowSelectedOptions(true).
		Show()

	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

