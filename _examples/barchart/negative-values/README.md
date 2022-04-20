# barchart/negative-values

![Animation](animation.svg)

```go
package main

import (
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

	pterm.Info.Println("Chart example with negative only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}

```
