### barchart/demo

![Animation](https://vhs.charm.sh/vhs-iJAm2JfQl3Y6BEHbigJLl.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// One deliberately longer label shows how the horizontal chart
	// aligns all bars to the widest label.
	bars := []pterm.Bar{
		{Label: "Go", Value: 5},
		{Label: "Rust", Value: 3},
		{Label: "TypeScript", Value: 7},
	}

	pterm.Info.Println("Chart example with positive values only (bars use the full chart area)")

	pterm.DefaultBarChart.WithBars(bars).Render()

	// The same data rendered horizontally.
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

	// WithHeight caps the chart at 5 rows; the bar values are scaled to fit.
	pterm.DefaultBarChart.WithBars(bars).WithHeight(5).Render()
}
```

</details>

### barchart/custom-width

![Animation](https://vhs.charm.sh/vhs-6GmiN2pa05uOn1wNnHqYQ3.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
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

	// In a horizontal chart, WithWidth limits how far the bars extend
	// to the right; the values are scaled to fit into 5 columns.
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithWidth(5).Render()
}
```

</details>

### barchart/default

![Animation](https://vhs.charm.sh/vhs-6CnMQBdI0NEbJHIeGu1eVH.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
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

	// By default the chart is drawn vertically, with each label under its bar.
	pterm.DefaultBarChart.WithBars(bars).Render()
}
```

</details>

### barchart/horizontal

![Animation](https://vhs.charm.sh/vhs-70ICgyL6YEO64KZ7uTmbI6.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
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

	// WithHorizontal draws one bar per line, left to right, with the
	// label in front of each bar.
	pterm.DefaultBarChart.WithBars(bars).WithHorizontal().Render()
}
```

</details>

### barchart/horizontal-show-value

![Animation](https://vhs.charm.sh/vhs-ZmO00oqwh8qoYzfOXK89S.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
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

	// WithShowValue prints each bar's numeric value at the end of the bar.
	pterm.DefaultBarChart.WithBars(barData).WithHorizontal().WithShowValue().Render()
}
```

</details>

### barchart/mixed-values

![Animation](https://vhs.charm.sh/vhs-7yhBY7Cax9nItpaGWGhVIN.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Quarterly results with both gains and losses. The longer label
	// shows how the horizontal chart aligns bars to the widest label.
	bars := []pterm.Bar{
		{Label: "Q1", Value: 2},
		{Label: "Q2", Value: -3},
		{Label: "Q3", Value: -2},
		{Label: "Q4", Value: 5},
		{Label: "Yearly Total", Value: 7},
	}

	pterm.DefaultSection.Println("Chart example with mixed values (the chart area is split between the positive and negative side, so bars get less space when the absolute values differ a lot)")

	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()

	// The same data rendered horizontally.
	pterm.DefaultBarChart.WithHorizontal().WithBars(bars).WithShowValue().Render()
}
```

</details>

### barchart/negative-values

![Animation](https://vhs.charm.sh/vhs-6e0BcJqwkyAr1fhLsP7DCD.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// With only negative values there is no positive side to draw, so the
	// bars still use the full chart area. The longer label shows how the
	// horizontal chart aligns bars to the widest label.
	negativeBars := pterm.Bars{
		{Label: "Q1", Value: -5},
		{Label: "Q2", Value: -3},
		{Label: "Yearly Total", Value: -7},
	}

	pterm.Info.Println("Chart example with negative values only (bars use the full chart area)")

	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()

	// The same data rendered horizontally.
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}
```

</details>

### barchart/show-value

![Animation](https://vhs.charm.sh/vhs-7tJExfAXHgGGO0FVnEsoPP.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
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

	// WithShowValue prints each bar's numeric value above the bar.
	pterm.DefaultBarChart.WithBars(bars).WithShowValue().Render()
}
```

</details>

