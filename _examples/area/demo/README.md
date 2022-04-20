# area/demo

![Animation](animation.svg)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("The previous text will stay in place, while the area updates.")
	pterm.Print("\n\n") // Add two new lines as spacer.

	area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.
	for i := 0; i < 10; i++ {
		str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(time.Now().Format("15:04:05"))).Srender() // Save current time in str.
		area.Update(str)                                                                                                // Update Area contents.
		time.Sleep(time.Second)
	}
	area.Stop()
}

```
