package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	negativeBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: -5,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: -3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: -7,
		},
	}

	fmt.Println("Chart example with negative only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(negativeBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).Render()
	fmt.Println("=====================================================================")
}
