package main

import "github.com/pterm/pterm"

func main() {
	bars := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// WithHeight caps the chart at 5 rows; the bar values are scaled to fit.
	pterm.DefaultBarChart.WithBars(bars).WithHeight(5).Render()
}
