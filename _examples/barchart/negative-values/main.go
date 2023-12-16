package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a set of bars with negative values.
	// Each bar is represented by a struct with a label and a value.
	negativeBars := pterm.Bars{
		{Label: "Bar 1", Value: -5},
		{Label: "Bar 2", Value: -3},
		{Label: "Longer Label", Value: -7},
	}

	// Print an informational message to the console.
	pterm.Info.Println("Chart example with negative only values (bars use 100% of chart area)")

	// Create a vertical bar chart with the defined bars.
	// The WithShowValue() option is used to display the value of each bar in the chart.
	// The Render() method is called to draw the chart.
	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()

	// Create a horizontal bar chart with the same bars.
	// The WithHorizontal() option is used to orient the chart horizontally.
	// The WithShowValue() option and Render() method are used in the same way as before.
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}
