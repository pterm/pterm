package main

import "github.com/pterm/pterm"

func main() {
	// The radius is measured in terminal rows. A small pie fits nicely
	// into dashboards next to other output.
	slices := pterm.PieSlices{
		{Label: "Reads", Value: 80},
		{Label: "Writes", Value: 20},
	}

	pterm.DefaultPieChart.WithRadius(3).WithSlices(slices).Render()
}
