# progressbar/custom-style

![Animation](https://vhs.charm.sh/vhs-6q2cOGHtVlL0B7qOi8jpZX.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Every visual part of the progressbar can be swapped out. Here the bar
	// gets a retro ASCII look; clearing BarPartialCharacters disables the
	// smooth block-glyph edge, which would clash with plain "#".
	p, _ := pterm.DefaultProgressbar.
		WithTotal(50).
		WithTitle("Installing").
		WithBarCharacter("#").
		WithLastCharacter("#").
		WithBarFiller("-").
		WithBarPartialCharacters(nil).
		WithTitleStyle(pterm.NewStyle(pterm.FgLightYellow)).
		WithBarStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithShowElapsedTime(false).
		Start()

	for i := 0; i < p.Total; i++ {
		p.Increment()
		time.Sleep(time.Millisecond * 60)
	}
}
```
