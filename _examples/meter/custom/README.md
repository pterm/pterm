# meter/custom

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// For a meter where full is good, flip the thresholds: Error while it
	// is low, Success once it is nearly there.
	pterm.Meter.
		WithThresholds(
			pterm.Threshold{From: 0, Tone: pterm.ToneError},
			pterm.Threshold{From: 0.5, Tone: pterm.ToneWarning},
			pterm.Threshold{From: 0.8, Tone: pterm.ToneSuccess},
		).
		WithWidth(30).
		WithItems(
			pterm.MeterItem{Label: "unit tests", Value: 0.91},
			pterm.MeterItem{Label: "integration tests", Value: 0.64},
			pterm.MeterItem{Label: "docs", Value: 0.32},
		).
		Render()

	pterm.Println()

	// A fixed tone replaces the thresholds.
	pterm.Meter.WithTone(pterm.TonePrimary).WithValue(0.6, "synced").Render()
}
```
