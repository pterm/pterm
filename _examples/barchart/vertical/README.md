# barchart/vertical

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Vertical charts draw columns over a baseline, exact to an eighth of
	// a line. One tone for all columns keeps a time series calm.
	pterm.BarChart.WithVertical().WithTone(pterm.TonePrimary).WithBars(pterm.Bars{
		{Label: "Mon", Value: 3.5},
		{Label: "Tue", Value: 6},
		{Label: "Wed", Value: 4.25},
		{Label: "Thu", Value: 8},
		{Label: "Fri", Value: 5.5},
	}).WithUnit("h").Render()

	pterm.Println()

	// Negative values hang below the baseline.
	pterm.BarChart.WithVertical().WithHeight(6).WithBars(pterm.Bars{
		{Label: "Q1", Value: 12.5},
		{Label: "Q2", Value: -4},
		{Label: "Q3", Value: -9.25},
		{Label: "Q4", Value: 18},
	}).Render()
}
```
