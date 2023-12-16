# barchart/custom-height

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a slice of Bar structs. Each struct represents a bar in the chart.
	// The Label field is the name of the bar and the Value field is the height of the bar.
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

	// Create and render a bar chart with the defined bars and a height of 5.
	// The WithBars method is used to set the bars of the chart.
	// The WithHeight method is used to set the height of the chart.
	// The Render method is used to display the chart in the terminal.
	pterm.DefaultBarChart.WithBars(bars).WithHeight(5).Render()
}

```
