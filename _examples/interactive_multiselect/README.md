### interactive_multiselect/custom-checkmarks

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/custom-checkmarks/animation.svg)

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

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	printer.Checkmark = &pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")}
	selectedOptions, _ := printer.Show()
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

</details>

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
	var options []string

	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}

```

</details>

