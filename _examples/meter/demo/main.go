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
