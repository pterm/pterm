### interactive_multiselect/demo

![Animation](https://vhs.charm.sh/vhs-5Ab3fT3ObEBGkEpU2Bp6u4.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options.
	var options []string

	// Populate the options slice with 100 options.
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Add 5 more options to the slice, indicating the availability of fuzzy searching.
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Use PTerm's interactive multiselect to present the options to the user and capture their selections.
	// The Show() method displays the options and waits for user input.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()

	// Print the selected options, highlighted in green.
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/custom-checkmarks

![Animation](https://vhs.charm.sh/vhs-1qHwd2WmvXUm3zoVAgFGPm.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Populate the options slice with 5 options
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Create a new interactive multiselect printer with the options
	// Disable the filter and define the checkmark symbols
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithCheckmark(&pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")})

	// Show the interactive multiselect and get the selected options
	selectedOptions, _ := printer.Show()

	// Print the selected options
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/custom-filter-placeholder

![Animation](https://vhs.charm.sh/vhs-3NtvedqCpBJW1DRt8jMg4q.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Generate 100 options and add them to the options slice
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Generate 5 additional options with a specific message and add them to the options slice
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Create a new interactive multiselect printer with custom filter placeholder text
	// The WithFilterInputPlaceholder method allows you to customize the placeholder text shown when filtering
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	// Use PTerm's interactive multiselect feature to present the options to the user and capture their selections
	// The Show() method displays the options and waits for the user's input
	selectedOptions, _ := printer.Show()

	// Display the selected options to the user with a green color for emphasis
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/custom-keys

![Animation](https://vhs.charm.sh/vhs-4OZy3vP4kGVWAAvhaZlp4G.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"atomicgo.dev/keyboard/keys"
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Populate the options slice with 5 options
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Create a new interactive multiselect printer with the options
	// Disable the filter and set the keys for confirming and selecting options
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space)

	// Show the interactive multiselect and get the selected options
	selectedOptions, _ := printer.Show()

	// Print the selected options
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

### interactive_multiselect/show-selected-options

![Animation](https://vhs.charm.sh/vhs-4XI4oApCAUAzEHQcPe5FFY.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options.
	var options []string

	// Populate the options slice with 100 options.
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Use PTerm's interactive multiselect to present the options to the user and capture their selections.
	// The Show() method displays the options and waits for user input.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithShowSelectedOptions(true).
		Show()

	// Print the selected options, highlighted in green.
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}
```

</details>

