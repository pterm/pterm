### interactive_multiselect/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/demo/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/custom-checkmarks/animation.svg)

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

### interactive_multiselect/custom-keys

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/custom-keys/animation.svg)

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

