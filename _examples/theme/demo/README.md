# theme/demo

![Animation](animation.gif)

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Every printer colors itself from the nine theme slots, so this list is
	// the whole palette pterm draws with.
	tones := []pterm.Tone{
		pterm.TonePrimary, pterm.ToneAccent, pterm.ToneSuccess,
		pterm.ToneWarning, pterm.ToneError, pterm.ToneInfo,
		pterm.ToneText, pterm.ToneMuted, pterm.ToneBorder,
	}

	for _, tone := range tones {
		swatch := pterm.NewStyle(pterm.ToneDefault).On(tone).Sprint("      ")
		fmt.Printf("%s %s %s\n", swatch, tone.Bold().Sprintf("%-8s", tone), pterm.ToneMuted.Sprint(tone.Color()))
	}

	fmt.Println()

	// Status messages pick their tone from the same slots.
	pterm.Success.Println("Built 12 packages")
	pterm.Info.Println("Cache is 2.3 GB")
	pterm.Warning.Println("3 dependencies are outdated")
	pterm.Error.Println("Could not reach the registry")
}
```
