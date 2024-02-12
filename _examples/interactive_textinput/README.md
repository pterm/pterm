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
	// Create an interactive text input with single line input mode and show it
	result, _ := pterm.DefaultInteractiveTextInput.Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's answer with an info prefix
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_textinput/default-value

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/default-value/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Create an interactive text input with single line input mode and show it
	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("Some default value").Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's answer with an info prefix
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
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	result, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	// Print the user's input prefixed with an informational message.
	// The '%s' placeholder is replaced with the user's input.
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
	// Create an interactive text input with a mask for password input
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")

	// Show the password input prompt and store the result
	result, _ := passwordInput.Show("Enter your password")

	// Get the default logger from PTerm
	logger := pterm.DefaultLogger

	// Log the received password (masked)
	// Note: In a real-world application, you should never log passwords
	logger.Info("Password received", logger.Args("password", result))
}

```

</details>

