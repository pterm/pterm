# interactive_confirm/timeout

![Animation](animation.svg)

```go
package main

import (
	"time"
	
	"github.com/pterm/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveConfirm.WithTimeout(time.Second * 3).Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

```
