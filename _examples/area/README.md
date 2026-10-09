### area/demo

![Animation](https://vhs.charm.sh/vhs-4yaFqOuIaXiKpeFNvdnh7T.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// A live clock: only the area redraws, anything printed before it stays put.
	pterm.Info.Println("The previous text will stay in place, while the area updates.")
	pterm.Print("\n\n")

	area, _ := pterm.DefaultArea.WithCenter().Start()

	for range 10 {
		// Render the current time as big letters into a string, then swap it
		// into the area. Srender returns the output instead of printing it.
		str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender()
		area.Update(str)
		time.Sleep(time.Second)
	}

	area.Stop()
}
```

</details>

### area/center

![Animation](https://vhs.charm.sh/vhs-2x22TSH9shIZed56kRoF5Z.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// WithCenter horizontally centers the area's content in the terminal.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintfln("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}
```

</details>

### area/default

![Animation](https://vhs.charm.sh/vhs-1kGxZ5xpByOoQdxUeQ0KaT.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintfln("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}
```

</details>

### area/dynamic-chart

![Animation](https://vhs.charm.sh/vhs-6oEUx93vsoXr8zY4qDjSTp.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### area/fullscreen

![Animation](https://vhs.charm.sh/vhs-3V7I9NzBSGRahRPGKGNOKo.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Fullscreen clears the terminal and gives the area the whole screen.
	area, _ := pterm.DefaultArea.WithFullscreen().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	// Stop clears the area and restores the terminal.
	area.Stop()
}
```

</details>

### area/fullscreen-center

![Animation](https://vhs.charm.sh/vhs-Gi75JoygwXMXaFkJLn4z8.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Fullscreen takes over the whole terminal; WithCenter centers the
	// content within it.
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()

	// Each Update redraws the area in place instead of appending new lines.
	for i := range 5 {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	// Stop clears the area and restores the terminal.
	area.Stop()
}
```

</details>

