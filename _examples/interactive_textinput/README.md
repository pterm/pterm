### interactive_textinput/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_textinput/multi-line

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/multi-line/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show() // Text input with multi line enabled
	pterm.Println()                                                       // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_textinput/password

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/password/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Enter your password")

	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
}

```

</details>

