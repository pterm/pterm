### barchart/custom-height

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/custom-height/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithHeight(5).Render()
}

```

</details>

### barchart/custom-width

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/custom-width/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithHorizontal().WithWidth(5).Render()
}

```

</details>

### barchart/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).Render()
}

```

</details>

### barchart/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	positiveBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
		},
	}

	pterm.Info.Println("Chart example with positive only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(positiveBars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(positiveBars).Render()
}

```

</details>

### barchart/horizontal

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/horizontal/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithHorizontal().Render()
}

```

</details>

### barchart/horizontal-show-value

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/horizontal-show-value/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithHorizontal().WithShowValue().Render()
}

```

</details>

### barchart/mixed-values

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/mixed-values/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	mixedBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 2,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: -3,
		},
		pterm.Bar{
			Label: "Bar 3",
			Value: -2,
		},
		pterm.Bar{
			Label: "Bar 4",
			Value: 5,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
		},
	}

	pterm.DefaultSection.Println("Chart example with mixed values (note screen space usage in case when ABSOLUTE values of negative and positive parts are differ too much)")
	_ = pterm.DefaultBarChart.WithBars(mixedBars).WithShowValue().Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(mixedBars).WithShowValue().Render()
}

```

</details>

### barchart/negative-values

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/negative-values/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	negativeBars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: -5,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: -3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: -7,
		},
	}

	pterm.Info.Println("Chart example with negative only values (bars use 100% of chart area)")
	_ = pterm.DefaultBarChart.WithBars(negativeBars).WithShowValue().Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(negativeBars).WithShowValue().Render()
}

```

</details>

### barchart/show-value

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/show-value/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBarChart.WithBars([]pterm.Bar{
		{Label: "A", Value: 10},
		{Label: "B", Value: 20},
		{Label: "C", Value: 30},
		{Label: "D", Value: 40},
		{Label: "E", Value: 50},
		{Label: "F", Value: 40},
		{Label: "G", Value: 30},
		{Label: "H", Value: 20},
		{Label: "I", Value: 10},
	}).WithShowValue().Render()
}

```

</details>

