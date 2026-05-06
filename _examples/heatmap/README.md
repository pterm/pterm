### heatmap/demo

![Animation](https://vhs.charm.sh/vhs-5Gg9cnXkTW44HZiZqIKaM1.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap. Each sub-array represents a row in the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the labels for the X and Y axes of the heatmap.
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Create a heatmap with the defined data and axis labels, and enable RGB colors.
	// Then render the heatmap.
	pterm.DefaultHeatmap.WithAxisData(headerData).WithData(data).WithEnableRGB().Render()
}
```

</details>

### heatmap/custom_colors

![Animation](https://vhs.charm.sh/vhs-31y7AvLKclklj1MfSXW9wq.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message
	pterm.Info.Println("The following table has no rgb (supported by every terminal), no axis data and a legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options, and render it
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(headerData).
		WithLegend(false).
		WithColors(pterm.BgBlue, pterm.BgRed, pterm.BgGreen, pterm.BgYellow).
		WithLegend().
		Render()
}
```

</details>

### heatmap/custom_legend

![Animation](https://vhs.charm.sh/vhs-1R4ftXBA6ji3DRObytiIPh.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the header data for the heatmap
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a custom legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options
	// Options are chained in a single line for simplicity
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(headerData).
		WithEnableRGB().
		WithLegendLabel("custom").
		WithLegendOnlyColoredCells().
		Render() // Render the heatmap
}
```

</details>

### heatmap/custom_rgb

![Animation](https://vhs.charm.sh/vhs-5gD43rIoUbSLhOoDXOsX4P.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap.
	axisLabels := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message.
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a legend.")
	pterm.Println()

	// Define the color range for the heatmap.
	rgbRange := []pterm.RGB{
		pterm.NewRGB(0, 0, 255),
		pterm.NewRGB(255, 0, 0),
		pterm.NewRGB(0, 255, 0),
		pterm.NewRGB(255, 255, 0),
	}

	// Create and render the heatmap.
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(axisLabels).
		WithEnableRGB().
		WithRGBRange(rgbRange...).
		Render()
}
```

</details>

### heatmap/no_grid

![Animation](https://vhs.charm.sh/vhs-GCMu734z7KipKZyJkJAKQ.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis data for the heatmap.
	axisData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message.
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options, then render it.
	pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(axisData).WithEnableRGB().WithLegend().WithGrid(false).Render()
}
```

</details>

### heatmap/separated

![Animation](https://vhs.charm.sh/vhs-7tVQlAJIY38qZzmqnxXWTB.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap.
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message.
	pterm.Info.Println("The following table has no rgb (supported by every terminal), no axis data and no legend.")
	pterm.Println()

	// Create the heatmap with the specified data and options, and render it.
	pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithLegend(false).Render()
}
```

</details>

