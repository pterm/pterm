package main

import "github.com/pterm/pterm"

func main() {
	// The first slice starts at twelve o'clock and the rest follow
	// clockwise. Values don't need to add up to 100, PTerm works out
	// each slice's share of the total.
	slices := pterm.PieSlices{
		{Label: "Go", Value: 45},
		{Label: "Rust", Value: 25},
		{Label: "TypeScript", Value: 20},
		{Label: "Other", Value: 10},
	}

	pterm.DefaultPieChart.WithSlices(slices).Render()
}
