# heatmap/values

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Response times in milliseconds per region and hour.
	data := pterm.HeatmapData{
		{42, 45, 51, 88, 120, 96},
		{38, 40, 47, 61, 75, 70},
		{110, 128, 142, 190, 240, 205},
	}

	axis := pterm.HeatmapAxis{
		XAxis: []string{"9h", "10h", "11h", "12h", "13h", "14h"},
		YAxis: []string{"eu", "us", "ap"},
	}

	// The values themselves, colored from Success to Warning to Error,
	// because high response times are bad.
	pterm.Heatmap.
		WithData(data).
		WithAxisData(axis).
		WithShowValues().
		WithGradient(pterm.TrafficTones()...).
		WithLegendLabel("ms").
		Render()
}
```
