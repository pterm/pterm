package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Start a new fullscreen centered area.
	// This area will be used to display the bar chart.
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()
	// Ensure the area stops updating when we're done.
	defer area.Stop()

	// Loop to update the bar chart 10 times.
	for i := 0; i < 10; i++ {
		// Create a new bar chart with dynamic bars.
		// The bars will change based on the current iteration.
		barchart := pterm.DefaultBarChart.WithBars(dynamicBars(i))
		// Render the bar chart to a string.
		// This string will be used to update the area.
		content, _ := barchart.Srender()
		// Update the area with the new bar chart.
		area.Update(content)
		// Wait for half a second before the next update.
		time.Sleep(500 * time.Millisecond)
	}
}

// dynamicBars generates a set of bars for the bar chart.
// The bars will change based on the current iteration.
func dynamicBars(i int) pterm.Bars {
	return pterm.Bars{
		{Label: "A", Value: 10},     // A static bar.
		{Label: "B", Value: 20 * i}, // A bar that grows with each iteration.
		{Label: "C", Value: 30},     // Another static bar.
		{Label: "D", Value: 40 + i}, // A bar that grows slowly with each iteration.
	}
}
