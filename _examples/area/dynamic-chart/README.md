# area/dynamic-chart

![Animation](https://vhs.charm.sh/vhs-6oEUx93vsoXr8zY4qDjSTp.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A live bar chart: render the chart to a string each tick and let the
	// area redraw it in place. Fullscreen gives the area the whole terminal.
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()
	defer area.Stop()

	for i := range 10 {
		barchart := pterm.DefaultBarChart.WithBars(dynamicBars(i))
		content, _ := barchart.Srender()
		area.Update(content)
		time.Sleep(500 * time.Millisecond)
	}
}

// dynamicBars returns the chart data for the given tick. B and D grow over
// time, A and C stay constant.
func dynamicBars(i int) pterm.Bars {
	return pterm.Bars{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20 * i},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40 + i},
	}
}
```
