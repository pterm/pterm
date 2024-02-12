### area/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/demo/animation.svg)

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
	// Print an informational message using PTerm's Info printer.
	// This message will stay in place while the area updates.
	pterm.Info.Println("The previous text will stay in place, while the area updates.")

	// Print two new lines as spacer.
	pterm.Print("\n\n")

	// Start the Area printer from PTerm's DefaultArea, with the Center option.
	// The Area printer allows us to update a specific area of the console output.
	// The returned 'area' object is used to control the area updates.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Loop 10 times to update the area with the current time.
	for i := 0; i < 10; i++ {
		// Get the current time, format it as "15:04:05" (hour:minute:second), and convert it to a string.
		// Then, create a BigText from the time string using PTerm's DefaultBigText and putils NewLettersFromString.
		// The Srender() function is used to save the BigText as a string.
		str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender()

		// Update the Area contents with the current time string.
		area.Update(str)

		// Sleep for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the Area printer after all updates are done.
	area.Stop()
}

```

</details>

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
	// Start a new default area in the center of the terminal.
	// The Start() function returns the created area and an error.
	area, _ := pterm.DefaultArea.WithCenter().Start()

	// Loop 5 times to simulate a dynamic update.
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count.
		// The Sprintf function is used to format the string.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second to simulate a time-consuming task.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
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
	// Start a new default area and get a reference to it.
	// The second return value is an error which is ignored here.
	area, _ := pterm.DefaultArea.Start()

	// Loop 5 times
	for i := 0; i < 5; i++ {
		// Update the content of the area dynamically.
		// Here we're just displaying the current count.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
	// This will clean up and free resources used by the area.
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
	// Start a new fullscreen area. This will return an area instance and an error.
	// The underscore (_) is used to ignore the error.
	area, _ := pterm.DefaultArea.WithFullscreen().Start()

	// Loop 5 times to update the area content.
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count.
		// The Sprintf function is used to format the string.
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second before the next update.
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done.
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
	// Initialize a new PTerm area with fullscreen and center options
	// The Start() function returns the created area and an error (ignored here)
	area, _ := pterm.DefaultArea.WithFullscreen().WithCenter().Start()

	// Loop 5 times to demonstrate dynamic content update
	for i := 0; i < 5; i++ {
		// Update the content of the area with the current count
		// The Sprintf function is used to format the string with the count
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))

		// Pause for a second
		time.Sleep(time.Second)
	}

	// Stop the area after all updates are done
	// This will clear the area and return the terminal to its normal state
	area.Stop()
}

```

</details>

