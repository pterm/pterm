package main

import "github.com/pterm/pterm"

func main() {
	// Show the raw values next to the percentages, handy for things
	// like disk usage where the absolute numbers matter too.
	slices := pterm.PieSlices{
		{Label: "Used (GB)", Value: 312.5},
		{Label: "Free (GB)", Value: 187.5},
	}

	pterm.DefaultPieChart.WithShowValue().WithSlices(slices).Render()
}
