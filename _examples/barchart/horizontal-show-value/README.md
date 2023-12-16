# barchart/horizontal-show-value

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart
	barData := []pterm.Bar{
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

	// Create a bar chart with the defined data
	// The chart is horizontal and displays the value of each bar
	// The Render() function is called to display the chart
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithShowValue().Render()
}

```
