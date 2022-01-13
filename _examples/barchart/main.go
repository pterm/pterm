package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	positiveBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
		},
	}

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

	mixedBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 2,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: -3,
		},
		pterm.Bar{
			Label: "Bar 3",
			Value: -2,
		},
		pterm.Bar{
			Label: "Bar 4",
			Value: 5,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
		},
	}

	fmt.Println("Chart example with positive only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(positiveBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(positiveBars).Render()
	fmt.Println("=====================================================================")

	fmt.Println("Chart example with negative only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(negativeBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).Render()
	fmt.Println("=====================================================================")

	fmt.Println("Chart example with mixed values (positive and negative bars use by 50% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(mixedBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(mixedBars).Render()
	fmt.Println("=====================================================================")
}
