# interactive_textinput/multi-line

![Animation](animation.svg)

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
