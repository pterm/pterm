# bigtext/colored

![Animation](https://vhs.charm.sh/vhs-6OVbbx3QPZS9Rz75SwAzmR.gif)

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Each LettersFromStringWithStyle call gets its own color, so parts of
	// the big text can be highlighted independently.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()
}
```
