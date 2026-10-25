package main

import "github.com/pterm/pterm"

func main() {
	// Give a slice its own style to pin its color. Slices without a
	// style keep cycling through the printer's default colors.
	slices := pterm.PieSlices{
		{Label: "Healthy", Value: 42, Style: pterm.NewStyle(pterm.FgGreen)},
		{Label: "Degraded", Value: 6, Style: pterm.NewStyle(pterm.FgYellow)},
		{Label: "Down", Value: 2, Style: pterm.NewStyle(pterm.FgRed)},
	}

	pterm.DefaultPieChart.WithSlices(slices).Render()

	// Or swap the whole palette at once.
	pterm.DefaultPieChart.
		WithColors([]pterm.Color{pterm.FgLightBlue, pterm.FgBlue, pterm.FgCyan}).
		WithSlices(pterm.PieSlices{
			{Label: "eu-west", Value: 5},
			{Label: "us-east", Value: 3},
			{Label: "ap-south", Value: 2},
		}).
		Render()
}
