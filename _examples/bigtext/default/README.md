# bigtext/default

![Animation](https://vhs.charm.sh/vhs-3AzGuVlmJ9yL1ADI8F1IV2.gif)

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// BigText takes Letters rather than a plain string; putils converts one
	// into the other.
	letters := putils.LettersFromString("PTerm")

	pterm.DefaultBigText.WithLetters(letters).Render()
}
```
