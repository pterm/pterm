package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a set of bars for the chart.
	// Each bar has a label and a value.
	bars := []pterm.Bar{
		{Label: "Bar 1", Value: 2},
		{Label: "Bar 2", Value: -3},
		{Label: "Bar 3", Value: -2},
		{Label: "Bar 4", Value: 5},
		{Label: "Longer Label", Value: 7},
	}

	// Print a section header.
	// This is useful for separating different parts of the output.
	pterm.DefaultSection.Println("Chart example with mixed values (note screen space usage in case when ABSOLUTE values of negative and positive parts are differ too much)")

	// Create a bar chart with the defined bars.
	// The chart will display the value of each bar.
	// The Render() function is called to display the chart.
	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()

	// Create a horizontal bar chart with the same bars.
	// The chart will display the value of each bar.
	// The Render() function is called to display the chart.
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).WithShowValue().Render()
}
