### area/center

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/center/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithCenter().Start()

	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

```

</details>

### area/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/default/animation.svg)

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

	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

```

</details>

### area/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("The previous text will stay in place, while the area updates.")
	pterm.Print("\n\n") // Add two new lines as spacer.

	area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.
	for i := 0; i < 10; i++ {
		str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(time.Now().Format("15:04:05"))).Srender() // Save current time in str.
		area.Update(str)                                                                                                // Update Area contents.
		time.Sleep(time.Second)
	}
	area.Stop()
}

```

</details>

### area/dynamic-chart

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/dynamic-chart/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
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

```

</details>

### area/fullscreen

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/fullscreen/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithFullscreen().Start()

	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

```

</details>

### area/fullscreen-center

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/fullscreen-center/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()

	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}

```

</details>

