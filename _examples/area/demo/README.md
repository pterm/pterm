# area/demo

![Animation](https://vhs.charm.sh/vhs-4yaFqOuIaXiKpeFNvdnh7T.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// A live clock: only the area redraws, anything printed before it stays put.
	pterm.Info.Println("The previous text will stay in place, while the area updates.")
	pterm.Print("\n\n")

	area, _ := pterm.DefaultArea.WithCenter().Start()

	for range 10 {
		// Render the current time as big letters into a string, then swap it
		// into the area. Srender returns the output instead of printing it.
		str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender()
		area.Update(str)
		time.Sleep(time.Second)
	}

	area.Stop()
}
```
