# barchart

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	bars := pterm.Bars{
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

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
}

```
