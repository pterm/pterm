# badge/chip

![Animation](animation.gif)

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Chips are the quiet variant: the text in the badge's tone, between
	// brackets. They work well for tags next to list items.
	chip := pterm.Badge.Chip()

	fmt.Println("api-server ", chip.WithTone(pterm.ToneSuccess).Sprint("running"))
	fmt.Println("worker     ", chip.WithTone(pterm.ToneWarning).Sprint("restarting"))
	fmt.Println("scheduler  ", chip.WithTone(pterm.ToneMuted).Sprint("stopped"))
}
```
