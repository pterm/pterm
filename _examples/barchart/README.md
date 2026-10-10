### barchart/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### barchart/custom

![Animation](custom/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	usage := pterm.Bars{
		{Label: "/", Value: 0.82},
		{Label: "/home", Value: 0.47},
		// A tone on a single bar wins over the chart's tones.
		{Label: "/var", Value: 0.96, Tone: pterm.ToneError},
	}

	// A fixed range makes bars comparable between charts, a custom format
	// prints the values as percentages, and the width includes labels and
	// values.
	pterm.BarChart.
		WithBars(usage).
		WithRange(0, 1).
		WithTone(pterm.ToneSuccess).
		WithFormat(func(v float64) string { return fmt.Sprintf("%.0f%%", v*100) }).
		WithWidth(50).
		Render()

	pterm.Println()

	// Without track and values, only the bars remain.
	pterm.BarChart.WithBars(usage).WithRange(0, 1).WithShowTrack(false).WithShowValues(false).WithWidth(40).Render()
}
```

</details>

### barchart/grouped

![Animation](grouped/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Each group holds one value per series, in the order of WithSeries.
	// The series names make up the legend under the chart.
	pterm.BarChart.
		WithSeries("2025", "2026").
		WithGroups(pterm.BarGroups{
			{Label: "Q1", Values: []float64{42, 51}},
			{Label: "Q2", Values: []float64{38, 57}},
			{Label: "Q3", Values: []float64{45, 63.5}},
		}).
		WithUnit("k").
		Render()
}
```

</details>

### barchart/negative

![Animation](negative/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Losses grow to the left of the zero axis, gains to the right.
	pterm.BarChart.WithBars(pterm.Bars{
		{Label: "Q1", Value: 12.5},
		{Label: "Q2", Value: -4},
		{Label: "Q3", Value: -9.25},
		{Label: "Q4", Value: 18},
	}).WithUnit("%").Render()
}
```

</details>

### barchart/stacked

![Animation](stacked/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Stacked bars show how a total splits up. The value on the right is
	// the total of each bar.
	pterm.BarChart.
		WithSeries("cpu", "io", "network").
		WithStacked().
		WithUnit(" ms").
		WithGroups(pterm.BarGroups{
			{Label: "build", Values: []float64{120, 45, 12}},
			{Label: "test", Values: []float64{210, 80, 5}},
			{Label: "deploy", Values: []float64{30, 20, 140}},
		}).
		Render()
}
```

</details>

### barchart/vertical

![Animation](vertical/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Vertical charts draw columns over a baseline, exact to an eighth of
	// a line. One tone for all columns keeps a time series calm.
	pterm.BarChart.WithVertical().WithTone(pterm.TonePrimary).WithBars(pterm.Bars{
		{Label: "Mon", Value: 3.5},
		{Label: "Tue", Value: 6},
		{Label: "Wed", Value: 4.25},
		{Label: "Thu", Value: 8},
		{Label: "Fri", Value: 5.5},
	}).WithUnit("h").Render()

	pterm.Println()

	// Negative values hang below the baseline.
	pterm.BarChart.WithVertical().WithHeight(6).WithBars(pterm.Bars{
		{Label: "Q1", Value: 12.5},
		{Label: "Q2", Value: -4},
		{Label: "Q3", Value: -9.25},
		{Label: "Q4", Value: 18},
	}).Render()
}
```

</details>

