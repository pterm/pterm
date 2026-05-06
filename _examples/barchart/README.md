### barchart/demo

![Animation](https://vhs.charm.sh/vhs-48hC9kutzQcqurPjdprkey.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the bars for the chart
	bars := []pterm.Bar{
		{Label: "Bar 1", Value: 5},
		{Label: "Bar 2", Value: 3},
		{Label: "Longer Label", Value: 7},
	}

	// Print an informational message
	pterm.Info.Println("Chart example with positive only values (bars use 100% of chart area)")

	// Create a bar chart with the defined bars and render it
	// The DefaultBarChart is used as a base, and the bars are added with the WithBars option
	// The Render function is then called to display the chart
	pterm.DefaultBarChart.WithBars(bars).Render()

	// Create a horizontal bar chart with the defined bars and render it
	// The DefaultBarChart is used as a base, the chart is made horizontal with the WithHorizontal option, and the bars are added with the WithBars option
	// The Render function is then called to display the chart
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
}
```

</details>

### barchart/custom-height

![Animation](https://vhs.charm.sh/vhs-5vLTJ082V7bklOqcjLFHIR.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a slice of Bar structs. Each struct represents a bar in the chart.
	// The Label field is the name of the bar and the Value field is the height of the bar.
	bars := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Create and render a bar chart with the defined bars and a height of 5.
	// The WithBars method is used to set the bars of the chart.
	// The WithHeight method is used to set the height of the chart.
	// The Render method is used to display the chart in the terminal.
	pterm.DefaultBarChart.WithBars(bars).WithHeight(5).Render()
}
```

</details>

### barchart/custom-width

![Animation](https://vhs.charm.sh/vhs-471reYokQ5iNiKbyhQawvz.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart
	barData := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Create a bar chart with the defined data
	// The chart is horizontal and has a width of 5
	// The Render() function is called to display the chart
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithWidth(5).Render()
}
```

</details>

### barchart/default

![Animation](https://vhs.charm.sh/vhs-1l29HBVXZEKqB5fR20U7yg.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart. Each bar is represented by a `pterm.Bar` struct.
	// The `Label` field represents the label of the bar, and the `Value` field represents the value of the bar.
	bars := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Use the `DefaultBarChart` from the `pterm` package to create a bar chart.
	// The `WithBars` method is used to set the bars of the chart.
	// The `Render` method is used to display the chart.
	pterm.DefaultBarChart.WithBars(bars).Render()
}
```

</details>

### barchart/horizontal

![Animation](https://vhs.charm.sh/vhs-42CW1xtORKZ1ZAg1r0Cv5X.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart
	bars := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Create a bar chart with the defined data
	// The chart is displayed horizontally
	// The Render() function is called to display the chart
	pterm.DefaultBarChart.WithBars(bars).WithHorizontal().Render()
}
```

</details>

### barchart/horizontal-show-value

![Animation](https://vhs.charm.sh/vhs-3iDMaVYU1VnylxmagojVzt.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the bar chart
	barData := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Create a bar chart with the defined data
	// The chart is horizontal and displays the value of each bar
	// The Render() function is called to display the chart
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithShowValue().Render()
}
```

</details>

### barchart/mixed-values

![Animation](https://vhs.charm.sh/vhs-54klL5UuBaS89PN935LwvX.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a set of bars for the chart.
	// Each bar has a label and a value.
	bars := []pterm.Bar{
		{Label: "Bar 1", Value: 2},
		{Label: "Bar 2", Value: -3},
		{Label: "Bar 3", Value: -2},
		{Label: "Bar 4", Value: 5},
		{Label: "Longer Label", Value: 7},
	}

	// Print a section header.
	// This is useful for separating different parts of the output.
	pterm.DefaultSection.Println("Chart example with mixed values (note screen space usage in case when ABSOLUTE values of negative and positive parts are differ too much)")

	// Create a bar chart with the defined bars.
	// The chart will display the value of each bar.
	// The Render() function is called to display the chart.
	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()

	// Create a horizontal bar chart with the same bars.
	// The chart will display the value of each bar.
	// The Render() function is called to display the chart.
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).WithShowValue().Render()
}
```

</details>

### barchart/negative-values

![Animation](https://vhs.charm.sh/vhs-4cB5U3tsHOwMpArQkPR2t2.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a set of bars with negative values.
	// Each bar is represented by a struct with a label and a value.
	negativeBars := pterm.Bars{
		{Label: "Bar 1", Value: -5},
		{Label: "Bar 2", Value: -3},
		{Label: "Longer Label", Value: -7},
	}

	// Print an informational message to the console.
	pterm.Info.Println("Chart example with negative only values (bars use 100% of chart area)")

	// Create a vertical bar chart with the defined bars.
	// The WithShowValue() option is used to display the value of each bar in the chart.
	// The Render() method is called to draw the chart.
	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()

	// Create a horizontal bar chart with the same bars.
	// The WithHorizontal() option is used to orient the chart horizontally.
	// The WithShowValue() option and Render() method are used in the same way as before.
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}
```

</details>

### barchart/show-value

![Animation](https://vhs.charm.sh/vhs-Pur7MhOe8rCRcRX8LO4P6.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a slice of bars for the bar chart. Each bar is represented by a struct
	// with a Label and a Value. The Label is a string that represents the name of the bar,
	// and the Value is an integer that represents the height of the bar.
	bars := []pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}

	// Create a bar chart with the defined bars using the DefaultBarChart object from PTerm.
	// Chain the WithBars method to set the bars of the chart.
	// Chain the WithShowValue method to display the value of each bar on the chart.
	// Finally, call the Render method to display the chart.
	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()
}
```

</details>

