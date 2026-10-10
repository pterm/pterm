### heatmap/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Commits per weekday, in two-hour slots. Each row is one day, and the
	// colors run from Border through Primary to Accent.
	data := pterm.HeatmapData{
		{3, 6, 4, 7, 12, 24, 34, 32, 28, 12, 3, 6},
		{5, 3, 6, 4, 14, 26, 31, 34, 25, 14, 5, 3},
		{7, 5, 3, 6, 11, 28, 33, 31, 27, 11, 7, 5},
		{4, 7, 5, 3, 13, 25, 35, 33, 24, 13, 4, 7},
		{6, 4, 7, 5, 10, 27, 32, 35, 26, 10, 6, 4},
		{1, 4, 2, 5, 5, 6, 11, 9, 10, 5, 1, 4},
		{3, 1, 4, 2, 7, 8, 8, 11, 7, 7, 3, 1},
	}

	// One X label per column and one Y label per row.
	axis := pterm.HeatmapAxis{
		XAxis: []string{"00", "02", "04", "06", "08", "10", "12", "14", "16", "18", "20", "22"},
		YAxis: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
	}

	pterm.Heatmap.WithData(data).WithAxisData(axis).WithLegendLabel("commits").Render()
}
```

</details>

### heatmap/grid

![Animation](grid/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A correlation matrix: values from -1 to 1.
	data := pterm.HeatmapData{
		{1, 0.82, -0.4, 0.1},
		{0.82, 1, -0.25, 0.3},
		{-0.4, -0.25, 1, -0.7},
		{0.1, 0.3, -0.7, 1},
	}

	axis := pterm.HeatmapAxis{
		XAxis: []string{"cpu", "mem", "idle", "io"},
		YAxis: []string{"cpu", "mem", "idle", "io"},
	}

	// A fixed range keeps 0 in the middle of the gradient. The frame
	// implies the grid.
	pterm.Heatmap.
		WithData(data).
		WithAxisData(axis).
		WithRange(-1, 1).
		WithGradient(pterm.ToneError, pterm.ToneBorder, pterm.ToneInfo).
		WithShowValues().
		WithBoxed().
		Render()
}
```

</details>

### heatmap/values

![Animation](values/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

