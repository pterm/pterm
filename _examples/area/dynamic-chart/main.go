package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()
	defer area.Stop()

	for i := 0; i < 10; i++ {
		barchart := pterm.DefaultBarChart.WithBars(dynamicBars(i))
		content, _ := barchart.Srender()
		area.Update(content)
		time.Sleep(500 * time.Millisecond)
	}
}

func dynamicBars(i int) pterm.Bars {
	return pterm.Bars{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20 * i},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40 + i},
	}
}
