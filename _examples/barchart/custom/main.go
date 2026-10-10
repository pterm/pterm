package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	usage := pterm.Bars{
		{Label: "/", Value: 0.82},
		{Label: "/home", Value: 0.47},
		// A tone on a single bar wins over the chart's tones.
		{Label: "/var", Value: 0.96, Tone: pterm.ToneError},
	}

	// A fixed range makes bars comparable between charts, a custom format
	// prints the values as percentages, and the width includes labels and
	// values.
	pterm.BarChart.
		WithBars(usage).
		WithRange(0, 1).
		WithTone(pterm.ToneSuccess).
		WithFormat(func(v float64) string { return fmt.Sprintf("%.0f%%", v*100) }).
		WithWidth(50).
		Render()

	pterm.Println()

	// Without track and values, only the bars remain.
	pterm.BarChart.WithBars(usage).WithRange(0, 1).WithShowTrack(false).WithShowValues(false).WithWidth(40).Render()
}
