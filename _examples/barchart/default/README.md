# barchart/default

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart. Each bar is represented by a `pterm.Bar` struct.
	// The `Label` field represents the label of the bar, and the `Value` field represents the value of the bar.
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

	// Use the `DefaultBarChart` from the `pterm` package to create a bar chart.
	// The `WithBars` method is used to set the bars of the chart.
	// The `Render` method is used to display the chart.
	pterm.DefaultBarChart.WithBars(bars).Render()
}

```
