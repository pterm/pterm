package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// With only negative values there is no positive side to draw, so the
	// bars still use the full chart area. The longer label shows how the
	// horizontal chart aligns bars to the widest label.
	negativeBars := pterm.Bars{
		{Label: "Q1", Value: -5},
		{Label: "Q2", Value: -3},
		{Label: "Yearly Total", Value: -7},
	}

	pterm.Info.Println("Chart example with negative values only (bars use the full chart area)")

	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()

	// The same data rendered horizontally.
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}
