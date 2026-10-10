package main

import "github.com/pterm/pterm"

func main() {
	// Stacked bars show how a total splits up. The value on the right is
	// the total of each bar.
	pterm.BarChart.
		WithSeries("cpu", "io", "network").
		WithStacked().
		WithUnit(" ms").
		WithGroups(pterm.BarGroups{
			{Label: "build", Values: []float64{120, 45, 12}},
			{Label: "test", Values: []float64{210, 80, 5}},
			{Label: "deploy", Values: []float64{30, 20, 140}},
		}).
		Render()
}
