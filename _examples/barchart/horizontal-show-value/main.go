package main

import "github.com/pterm/pterm"

func main() {
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

	// WithShowValue prints each bar's numeric value at the end of the bar.
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithShowValue().Render()
}
