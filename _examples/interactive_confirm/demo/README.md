# interactive_confirm/demo

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Show an interactive confirmation dialog and get the result.
	result, _ := pterm.DefaultInteractiveConfirm.Show()

	// Print a blank line for better readability.
	pterm.Println()

	// Print the user's answer in a formatted way.
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

// boolToText converts a boolean value to a colored text.
// If the value is true, it returns a green "Yes".
// If the value is false, it returns a red "No".
func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

```
