package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the bars for the chart
	bars := []pterm.Bar{
		{Label: "Bar 1", Value: 5},
		{Label: "Bar 2", Value: 3},
		{Label: "Longer Label", Value: 7},
	}

	// Print an informational message
	pterm.Info.Println("Chart example with positive only values (bars use 100% of chart area)")

	// Create a bar chart with the defined bars and render it
	// The DefaultBarChart is used as a base, and the bars are added with the WithBars option
	// The Render function is then called to display the chart
	pterm.DefaultBarChart.WithBars(bars).Render()

	// Create a horizontal bar chart with the defined bars and render it
	// The DefaultBarChart is used as a base, the chart is made horizontal with the WithHorizontal option, and the bars are added with the WithBars option
	// The Render function is then called to display the chart
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
}
