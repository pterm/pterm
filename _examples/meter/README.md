### meter/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A single meter: the fill, the percentage and a muted label.
	pterm.Meter.WithValue(0.34, "of the monthly budget").Render()

	pterm.Println()

	// Fills turn from Success to Warning to Error as they get full.
	pterm.Meter.WithItems(
		pterm.MeterItem{Label: "cpu", Value: 0.42},
		pterm.MeterItem{Label: "memory", Value: 0.785},
		pterm.MeterItem{Label: "disk", Value: 0.93},
		pterm.MeterItem{Label: "swap", Value: 0.06},
	).Render()
}
```

</details>

### meter/custom

![Animation](custom/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

