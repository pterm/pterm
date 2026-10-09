# interactive_confirm/demo

![Animation](https://vhs.charm.sh/vhs-5fNTnAUGfbaRlaCGkWMBk3.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The confirm prompt accepts "y" and "n" as shortcuts. Pressing enter
	// answers with the default value, which is "no" unless changed with
	// WithDefaultValue.
	result, _ := pterm.DefaultInteractiveConfirm.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

// boolToText renders the answer as a colored "Yes" or "No".
func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}
```
