# area/demo

![Animation](animation.svg)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Print an informational message using PTerm's Info printer.
	// This message will stay in place while the area updates.
	pterm.Info.Println("The previous text will stay in place, while the area updates.")

	// Print two new lines as spacer.
	pterm.Print("\n\n")

	// Start the Area printer from PTerm's DefaultArea, with the Center option.
	// The Area printer allows us to update a specific area of the console output.
	// The returned 'area' object is used to control the area updates.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Loop 10 times to update the area with the current time.
	for i := 0; i < 10; i++ {
		// Get the current time, format it as "15:04:05" (hour:minute:second), and convert it to a string.
		// Then, create a BigText from the time string using PTerm's DefaultBigText and putils NewLettersFromString.
		// The Srender() function is used to save the BigText as a string.
		str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender()

		// Update the Area contents with the current time string.
		area.Update(str)

		// Sleep for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the Area printer after all updates are done.
	area.Stop()
}

```
