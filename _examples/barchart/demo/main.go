package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// One deliberately longer label shows how the horizontal chart
	// aligns all bars to the widest label.
	bars := []pterm.Bar{
		{Label: "Go", Value: 5},
		{Label: "Rust", Value: 3},
		{Label: "TypeScript", Value: 7},
	}

	pterm.Info.Println("Chart example with positive values only (bars use the full chart area)")

	pterm.DefaultBarChart.WithBars(bars).Render()

	// The same data rendered horizontally.
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
}
