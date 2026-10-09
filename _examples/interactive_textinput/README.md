### interactive_textinput/demo

![Animation](https://vhs.charm.sh/vhs-7xu04W1sjODqqV4xW1bQrY.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### interactive_textinput/default-value

![Animation](https://vhs.charm.sh/vhs-5TfbfjgGuvVPj5xOcB8nQZ.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The default value is shown as a pre-filled suggestion. Pressing enter
	// right away returns it; typing anything replaces it.
	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("Some default value").Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
```

</details>

### interactive_textinput/multi-line

![Animation](https://vhs.charm.sh/vhs-6BcAPqMvP68DHVX0tCfZg5.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### interactive_textinput/password

![Animation](https://vhs.charm.sh/vhs-7a9giWv4rH3696aO44jmgZ.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// WithMask echoes the given string instead of the typed characters,
	// which turns the text input into a password prompt.
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")

	result, _ := passwordInput.Show("Enter your password")

	// Never log passwords in a real application, this is just a demo.
	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
}
```

</details>

