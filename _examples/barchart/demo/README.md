# barchart/demo

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Weekly downloads per package manager. Each bar takes the next tone
	// of the theme's series cycle, and the values line up on the right.
	pterm.BarChart.WithBars(pterm.Bars{
		{Label: "npm", Value: 1284},
		{Label: "pnpm", Value: 932.5},
		{Label: "yarn", Value: 611},
		{Label: "bun", Value: 287},
	}).Render()

	pterm.Println()

	// Several series per label become a stacked bar with a legend.
	pterm.BarChart.
		WithSeries("compile", "test", "upload").
		WithStacked().
		WithUnit("s").
		WithGroups(pterm.BarGroups{
			{Label: "linux", Values: []float64{48, 31, 6}},
			{Label: "macos", Values: []float64{62, 35, 9}},
			{Label: "windows", Values: []float64{81, 44, 12}},
		}).
		Render()
}
```
