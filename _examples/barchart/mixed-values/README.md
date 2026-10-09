# barchart/mixed-values

![Animation](https://vhs.charm.sh/vhs-7yhBY7Cax9nItpaGWGhVIN.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Quarterly results with both gains and losses. The longer label
	// shows how the horizontal chart aligns bars to the widest label.
	bars := []pterm.Bar{
		{Label: "Q1", Value: 2},
		{Label: "Q2", Value: -3},
		{Label: "Q3", Value: -2},
		{Label: "Q4", Value: 5},
		{Label: "Yearly Total", Value: 7},
	}

	pterm.DefaultSection.Println("Chart example with mixed values (the chart area is split between the positive and negative side, so bars get less space when the absolute values differ a lot)")

	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()

	// The same data rendered horizontally.
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).WithShowValue().Render()
}
```
