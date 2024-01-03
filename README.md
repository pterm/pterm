<!--suppress HtmlDeprecatedAttribute -->

<h1 align="center">💻 PTerm | Pretty Terminal Printer</h1>
<p align="center">A modern Go framework to make beautiful CLIs</p>

<p align="center">

<a href="https://github.com/pterm/pterm/releases" style="text-decoration: none">
<img src="https://img.shields.io/github/v/release/pterm/pterm?style=flat-square" alt="Latest Release">
</a>

<a href="https://github.com/pterm/pterm/stargazers" style="text-decoration: none">
<img src="https://img.shields.io/github/stars/pterm/pterm.svg?style=flat-square" alt="Stars">
</a>

<a href="https://github.com/pterm/pterm/fork" style="text-decoration: none">
<img src="https://img.shields.io/github/forks/pterm/pterm.svg?style=flat-square" alt="Forks">
</a>

<a href="https://opensource.org/licenses/MIT" style="text-decoration: none">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<a href="https://codecov.io/gh/pterm/pterm" style="text-decoration: none">
<img src="https://img.shields.io/codecov/c/gh/pterm/pterm?color=magenta&logo=codecov&style=flat-square" alt="Downloads">
</a>

<a href="https://codecov.io/gh/pterm/pterm" style="text-decoration: none">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-28774-magenta?style=flat-square" alt="Forks"><!-- unittestcount:end -->
</a>

<br/>

<a href="https://github.com/pterm/pterm/releases" style="text-decoration: none">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

 <a href="https://marvin.ws/twitter">
        <img src="https://img.shields.io/badge/Twitter-%40MarvinJWendt-1DA1F2?logo=twitter&style=for-the-badge"/>
    </a>

<br/>
<br/>

<a href="https://github.com/pterm/pterm/tree/master/_examples/demo/demo" style="text-decoration: none">
<img src="https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/demo/animation.svg" alt="PTerm">
</a>
<p align="center"><a href="https://github.com/pterm/pterm/tree/master/_examples/demo/demo" >Show Demo Code</p></p>

</p>

---

<p align="center">
<strong><a href="https://pterm.sh">PTerm.sh</a></strong>
|
<strong><a href="#-installation">Installation</a></strong>
|
<strong><a href="https://docs.pterm.sh/getting-started">Getting Started</a></strong>
|
<strong><a href="https://docs.pterm.sh/">Documentation</a></strong>
|
<strong><a href="https://github.com/pterm/pterm/tree/master/_examples">Examples</a></strong>
|
<strong><a href="https://github.com/pterm/pterm/discussions?discussions_q=category%3AQ%26A">Q&A</a></strong>
|
<strong><a href="https://discord.gg/vE2dNkfAmF">Discord</a></strong>
</p>

---

## 📦 Installation

To make PTerm available in your project, you can run the following command.\
Make sure to run this command inside your project, when you're using go modules 😉

```sh
go get github.com/pterm/pterm
```

## ⭐ Main Features

| Feature          | Description                                         |
|------------------|-----------------------------------------------------|
| 🪀 Easy to use    | PTerm emphasizes ease of use, with [examples](#-examples) and consistent component design. |
| 🤹‍♀️ Cross-Platform | PTerm works on various OS and terminals, including `Windows CMD`, `macOS iTerm2`, and in CI systems like `GitHub Actions`. |
| 🧪 Well tested    | A high test coverage and <!-- unittestcount2:start -->`28774`<!-- unittestcount2:end --> automated tests ensure PTerm's reliability. |
| ✨ Consistent Colors | PTerm uses the [ANSI color scheme](https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit) for uniformity and supports `TrueColor` for advanced terminals. |
| 📚 Component system | PTerm's flexible `Printers` can be used individually or combined to generate beautiful console output. |
| 🛠 Configurable   | PTerm is ready to use without configuration but allows easy customization for unique terminal output. |
| ✏ Documentation  | Access comprehensive docs on [pkg.go.dev](https://pkg.go.dev/github.com/pterm/pterm#section-documentation) and view practical examples in the [examples section](#-examples). |

### Printers (Components)

<div align="center">

<!-- printers:start -->
| Feature | Feature | Feature | Feature | Feature |
| :-------: | :-------: | :-------: | :-------: | :-------: |
| Area <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/area) |Barchart <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/barchart) |Basictext <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/basictext) |Bigtext <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/bigtext) |Box <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/box) |
| Bulletlist <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/bulletlist) |Center <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/center) |Coloring <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/coloring) |Header <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/header) |Heatmap <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/heatmap) |
| Interactive confirm <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_confirm) |Interactive continue <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_continue) |Interactive multiselect <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_multiselect) |Interactive select <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_select) |Interactive textinput <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_textinput) |
| Logger <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/logger) |Multiple-live-printers <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/multiple-live-printers) |Panel <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/panel) |Paragraph <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/paragraph) |Prefix <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/prefix) |
| Progressbar <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/progressbar) |Section <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/section) |Slog <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/slog) |Spinner <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/spinner) |Style <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/style) |
| Table <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/table) |Theme <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/theme) |Tree <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/tree) | |  | 
<!-- printers:end -->

</div>

---

<div align="center">

### 🦸‍♂️ Sponsors

<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg" />

---

</div>

## 🧪 Examples

<p align="center">
<table>
<tbody>
<td align="center">
<img width="2000" height="0"><br>
<a href="https://github.com/pterm/pterm/tree/master/_examples">‼️ You can find all the examples, in a much better structure and their source code, in "_examples" ‼️</a><br>
<sub>Click on the link above to show the examples folder.</sub>
<img width="2000" height="0">
</td>
</tbody>
</table>
</p>

<!-- examples:start -->
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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/custom-height/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/custom-width/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/default/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/horizontal/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/horizontal-show-value/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/mixed-values/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/negative-values/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/show-value/animation.svg)

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

### basictext/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/basictext/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The DefaultBasicText is a basic text printer provided by PTerm.
	// It is used to print text without any special formatting.
	pterm.DefaultBasicText.Println("Default basic text printer.")

	// The DefaultBasicText can be used in any context that requires a TextPrinter.
	// Here, we're using it with the LightMagenta function to color a portion of the text.
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")

	// The DefaultBasicText is also useful for resolving progress bars and spinners.
}

```

</details>

### bigtext/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bigtext/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Create a large text with the LetterStyle from the standard theme.
	// This is useful for creating title screens.
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()

	// Create a large text with differently colored letters.
	// Here, the first letter 'P' is colored cyan and the rest 'Term' is colored light magenta.
	// This can be used to highlight specific parts of the text.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle()),
	).Render()

	// Create a large text with a specific RGB color.
	// This can be used when you need a specific color that is not available in the standard colors.
	// Here, the color is gold (RGB: 255, 215, 0).
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0)),
	).Render()
}

```

</details>

### bigtext/colored

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bigtext/colored/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Initialize a big text display with the letters "P" and "Term"
	// "P" is displayed in cyan and "Term" is displayed in light magenta
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render() // Render the big text to the terminal
}

```

</details>

### bigtext/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bigtext/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Define the text to be rendered
	var text = "PTerm"

	// Convert the text into a format suitable for PTerm
	var letters = putils.LettersFromString(text)

	// Render the text using PTerm's default big text style
	pterm.DefaultBigText.WithLetters(letters).Render()
}

```

</details>

### box/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print an informational message.
	pterm.Info.Println("This might not be rendered correctly on GitHub,\nbut it will work in a real terminal.\nThis is because GitHub does not use a monospaced font by default for SVGs")

	// Create three panels with text, some of them with titles.
	// The panels are created using the DefaultBox style.
	panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	panel3 := pterm.DefaultBox.WithTitle("bottom center title").WithTitleBottomCenter().Sprint("Duis aute irure\ndolor in reprehenderit\nin voluptate velit esse cillum\ndolore eu fugiat\nnulla pariatur.")

	// Combine the panels into a layout using the DefaultPanel style.
	// The layout is a 2D grid, with each row being an array of panels.
	// In this case, the first row contains panel1 and panel2, and the second row contains only panel3.
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: panel1}, {Data: panel2}},
		{{Data: panel3}},
	}).Srender()

	// Print the panels layout inside a box with a title.
	// The box is created using the DefaultBox style, with the title positioned at the bottom right.
	pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
}

```

</details>

### box/custom-padding

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/custom-padding/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a default box with custom padding options and print "Hello, World!" inside it.
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Hello, World!")
}

```

</details>

### box/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a default box with PTerm and print a message in it.
	// The DefaultBox.Println method automatically starts, prints the message, and stops the box.
	pterm.DefaultBox.Println("Hello, World!")
}

```

</details>

### box/title

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/title/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a default box with specified padding
	paddedBox := pterm.DefaultBox.WithLeftPadding(4).WithRightPadding(4).WithTopPadding(1).WithBottomPadding(1)

	// Define a title for the box
	title := pterm.LightRed("I'm a box!")

	// Create boxes with the title positioned differently and containing different content
	box1 := paddedBox.WithTitle(title).Sprint("Hello, World!\n      1")                         // Title at default position (top left)
	box2 := paddedBox.WithTitle(title).WithTitleTopCenter().Sprint("Hello, World!\n      2")    // Title at top center
	box3 := paddedBox.WithTitle(title).WithTitleTopRight().Sprint("Hello, World!\n      3")     // Title at top right
	box4 := paddedBox.WithTitle(title).WithTitleBottomRight().Sprint("Hello, World!\n      4")  // Title at bottom right
	box5 := paddedBox.WithTitle(title).WithTitleBottomCenter().Sprint("Hello, World!\n      5") // Title at bottom center
	box6 := paddedBox.WithTitle(title).WithTitleBottomLeft().Sprint("Hello, World!\n      6")   // Title at bottom left
	box7 := paddedBox.WithTitle(title).WithTitleTopLeft().Sprint("Hello, World!\n      7")      // Title at top left

	// Render the boxes in a panel layout
	pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{box1}, {box2}, {box3}},
		{{box4}, {box5}, {box6}},
		{{box7}},
	}).Render()
}

```

</details>

### bulletlist/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Define a list of bullet list items with different levels.
	bulletListItems := []pterm.BulletListItem{
		{Level: 0, Text: "Level 0"}, // Level 0 item
		{Level: 1, Text: "Level 1"}, // Level 1 item
		{Level: 2, Text: "Level 2"}, // Level 2 item
	}

	// Use the default bullet list style to render the list items.
	pterm.DefaultBulletList.WithItems(bulletListItems).Render()

	// Define a string with different levels of indentation.
	text := `0
 1
  2
   3`

	// Convert the indented string to a bullet list and render it.
	putils.BulletListFromString(text, " ").Render()
}

```

</details>

### bulletlist/customized

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/customized/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a list of bullet list items with different styles and levels.
	bulletListItems := []pterm.BulletListItem{
		{
			Level:       0,                            // Level 0 (top level)
			Text:        "Blue",                       // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgBlue), // Text color
			BulletStyle: pterm.NewStyle(pterm.FgRed),  // Bullet color
		},
		{
			Level:       1,                                  // Level 1 (sub-item)
			Text:        "Green",                            // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgGreen),      // Text color
			Bullet:      "-",                                // Custom bullet symbol
			BulletStyle: pterm.NewStyle(pterm.FgLightWhite), // Bullet color
		},
		{
			Level:       2,                              // Level 2 (sub-sub-item)
			Text:        "Cyan",                         // Text to display
			TextStyle:   pterm.NewStyle(pterm.FgCyan),   // Text color
			Bullet:      ">",                            // Custom bullet symbol
			BulletStyle: pterm.NewStyle(pterm.FgYellow), // Bullet color
		},
	}

	// Create a bullet list with the defined items and render it.
	pterm.DefaultBulletList.WithItems(bulletListItems).Render()
}

```

</details>

### center/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/center/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Print a block of text centered in the terminal
	pterm.DefaultCenter.Println("This text is centered!\nIt centers the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters and store in 's'
	s, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Srender()

	// Print the BigLetters 's' centered in the terminal
	pterm.DefaultCenter.Println(s)

	// Print each line of the text separately centered in the terminal
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("This text is centered!\nBut each line is\ncentered\nseparately")
}

```

</details>

### coloring/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a table with different foreground and background colors.
	pterm.DefaultTable.WithData([][]string{
		{pterm.FgBlack.Sprint("Black"), pterm.FgRed.Sprint("Red"), pterm.FgGreen.Sprint("Green"), pterm.FgYellow.Sprint("Yellow")},
		{"", pterm.FgLightRed.Sprint("Light Red"), pterm.FgLightGreen.Sprint("Light Green"), pterm.FgLightYellow.Sprint("Light Yellow")},
		{pterm.BgBlack.Sprint("Black"), pterm.BgRed.Sprint("Red"), pterm.BgGreen.Sprint("Green"), pterm.BgYellow.Sprint("Yellow")},
		{"", pterm.BgLightRed.Sprint("Light Red"), pterm.BgLightGreen.Sprint("Light Green"), pterm.BgLightYellow.Sprint("Light Yellow")},
		{pterm.FgBlue.Sprint("Blue"), pterm.FgMagenta.Sprint("Magenta"), pterm.FgCyan.Sprint("Cyan"), pterm.FgWhite.Sprint("White")},
		{pterm.FgLightBlue.Sprint("Light Blue"), pterm.FgLightMagenta.Sprint("Light Magenta"), pterm.FgLightCyan.Sprint("Light Cyan"), pterm.FgLightWhite.Sprint("Light White")},
		{pterm.BgBlue.Sprint("Blue"), pterm.BgMagenta.Sprint("Magenta"), pterm.BgCyan.Sprint("Cyan"), pterm.BgWhite.Sprint("White")},
		{pterm.BgLightBlue.Sprint("Light Blue"), pterm.BgLightMagenta.Sprint("Light Magenta"), pterm.BgLightCyan.Sprint("Light Cyan"), pterm.BgLightWhite.Sprint("Light White")},
	}).Render() // Render the table.

	pterm.Println()

	// Print words in different colors.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))

	pterm.Println()

	// Create a new style with a red background, light green foreground, and bold text.
	style := pterm.NewStyle(pterm.BgRed, pterm.FgLightGreen, pterm.Bold)
	// Print text using the created style.
	style.Println("This text uses a style and is bold and light green with a red background!")
}

```

</details>

### coloring/disable-output

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/disable-output/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Loop from 0 to 14
	for i := 0; i < 15; i++ {
		switch i {
		case 5:
			// At the 5th iteration, print a message and disable the output
			pterm.Info.Println("Disabled Output!")
			pterm.DisableOutput()
		case 10:
			// At the 10th iteration, enable the output and print a message
			pterm.EnableOutput()
			pterm.Info.Println("Enabled Output!")
		}

		// Print a progress message for each iteration
		pterm.Printf("Printing something... [%d/%d]\n", i, 15)
	}
}

```

</details>

### coloring/fade-colors

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/fade-colors/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Print an informational message.
	pterm.Info.Println("RGB colors only work in Terminals which support TrueColor.")

	// Define the start and end points for the color gradient.
	startColor := pterm.NewRGB(0, 255, 255) // Cyan
	endColor := pterm.NewRGB(255, 0, 255)   // Magenta

	// Get the terminal height to determine the gradient range.
	terminalHeight := pterm.GetTerminalHeight()

	// Loop over the range of the terminal height to create a color gradient.
	for i := 0; i < terminalHeight-2; i++ {
		// Calculate the fade factor for the current step in the gradient.
		fadeFactor := float32(i) / float32(terminalHeight-2)

		// Create a color that represents the current step in the gradient.
		currentColor := startColor.Fade(0, 1, fadeFactor, endColor)

		// Print a string with the current color.
		currentColor.Println("Hello, World!")
	}
}

```

</details>

### coloring/fade-colors-rgb-style

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/fade-colors-rgb-style/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	// Define RGB colors
	white := pterm.NewRGB(255, 255, 255)
	grey := pterm.NewRGB(128, 128, 128)
	black := pterm.NewRGB(0, 0, 0)
	red := pterm.NewRGB(255, 0, 0)
	purple := pterm.NewRGB(255, 0, 255)
	green := pterm.NewRGB(0, 255, 0)

	// Define strings to be printed
	str1 := "RGB colors only work in Terminals which support TrueColor."
	str2 := "The background and foreground colors can be customized individually."
	str3 := "Styles can also be applied. For example: Bold or Italic."

	// Print first string with color fading from white to purple
	printFadedString(str1, white, purple, grey, black)

	// Print second string with color fading from purple to red
	printFadedString(str2, black, purple, red, red)

	// Print third string with color fading from white to green and style changes
	printStyledString(str3, white, green, red, black)
}

// printFadedString prints a string with color fading effect
func printFadedString(str string, fgStart, fgEnd, bgStart, bgEnd pterm.RGB) {
	strs := strings.Split(str, "")
	var result string
	for i := 0; i < len(str); i++ {
		// Create a style with color fading effect
		style := pterm.NewRGBStyle(fgStart.Fade(0, float32(len(str)), float32(i), fgEnd), bgStart.Fade(0, float32(len(str)), float32(i), bgEnd))
		// Append styled letter to result string
		result += style.Sprint(strs[i])
	}
	pterm.Println(result)
}

// printStyledString prints a string with color fading and style changes
func printStyledString(str string, fgStart, fgEnd, bgStart, bgEnd pterm.RGB) {
	strs := strings.Split(str, "")
	var result string
	boldStr := strings.Split("Bold", "")
	italicStr := strings.Split("Italic", "")
	bold, italic := 0, 0
	for i := 0; i < len(str); i++ {
		// Create a style with color fading effect
		style := pterm.NewRGBStyle(fgStart.Fade(0, float32(len(str)), float32(i), fgEnd), bgStart.Fade(0, float32(len(str)), float32(i), bgEnd))
		// Check if the next letters are "Bold" or "Italic" and add the corresponding style
		if bold < len(boldStr) && i+len(boldStr)-bold <= len(strs) && strings.Join(strs[i:i+len(boldStr)-bold], "") == strings.Join(boldStr[bold:], "") {
			style = style.AddOptions(pterm.Bold)
			bold++
		} else if italic < len(italicStr) && i+len(italicStr)-italic < len(strs) && strings.Join(strs[i:i+len(italicStr)-italic], "") == strings.Join(italicStr[italic:], "") {
			style = style.AddOptions(pterm.Italic)
			italic++
		}
		// Append styled letter to result string
		result += style.Sprint(strs[i])
	}
	pterm.Println(result)
}

```

</details>

### coloring/fade-multiple-colors

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/fade-multiple-colors/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	// Define RGB values for gradient points.
	startColor := pterm.NewRGB(0, 255, 255)
	firstPoint := pterm.NewRGB(255, 0, 255)
	secondPoint := pterm.NewRGB(255, 0, 0)
	thirdPoint := pterm.NewRGB(0, 255, 0)
	endColor := pterm.NewRGB(255, 255, 255)

	// Define the string to be printed.
	str := "RGB colors only work in Terminals which support TrueColor."
	strs := strings.Split(str, "")

	// Initialize an empty string for the faded info.
	var fadeInfo string

	// Loop over the string length to create a gradient effect.
	for i := 0; i < len(str); i++ {
		// Append each character of the string with a faded color to the info string.
		fadeInfo += startColor.Fade(0, float32(len(str)), float32(i), firstPoint).Sprint(strs[i])
	}

	// Print the info string with gradient effect.
	pterm.Info.Println(fadeInfo)

	// Get the terminal height.
	terminalHeight := pterm.GetTerminalHeight()

	// Loop over the terminal height to print "Hello, World!" with a gradient effect.
	for i := 0; i < terminalHeight-2; i++ {
		// Print the string with a color that fades from startColor to endColor.
		startColor.Fade(0, float32(terminalHeight-2), float32(i), firstPoint, secondPoint, thirdPoint, endColor).Println("Hello, World!")
	}
}

```

</details>

### coloring/override-default-printers

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/override-default-printers/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a default error message with PTerm's built-in Error style.
	pterm.Error.Println("This is the default Error")

	// Override the default error prefix with a new text and style.
	pterm.Error.Prefix = pterm.Prefix{Text: "OVERRIDE", Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed)}

	// Print the error message again, this time with the overridden prefix.
	pterm.Error.Println("This is the default Error after the prefix was overridden")
}

```

</details>

### coloring/print-color-rgb

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/print-color-rgb/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a new RGB color with values 178, 44, 199.
	// This color will be used for the text.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")

	// Create a new RGB color with values 15, 199, 209.
	// This color will be used for the text.
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")

	// Create a new RGB color with values 201, 144, 30.
	// This color will be used for the background.
	// The 'true' argument indicates that the color is for the background.
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}

```

</details>

### coloring/print-color-rgb-style

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/coloring/print-color-rgb-style/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define RGB colors for foreground and background.
	foregroundRGB := pterm.RGB{R: 187, G: 80, B: 0}
	backgroundRGB := pterm.RGB{R: 0, G: 50, B: 123}

	// Create a new RGB style with the defined foreground and background colors.
	rgbStyle := pterm.NewRGBStyle(foregroundRGB, backgroundRGB)

	// Print a string with the custom RGB style.
	rgbStyle.Println("This text is not styled.")

	// Add the 'Bold' option to the RGB style and print a string with this style.
	rgbStyle.AddOptions(pterm.Bold).Println("This text is bold.")

	// Add the 'Italic' option to the RGB style and print a string with this style.
	rgbStyle.AddOptions(pterm.Italic).Println("This text is italic.")
}

```

</details>

### demo/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"flag"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// Speed the demo up, by setting this flag.
// Usefull for debugging.
// Example:
//
//	go run main.go -speedup
var speedup = flag.Bool("speedup", false, "Speed up the demo")
var skipIntro = flag.Bool("skip-intro", false, "Skips the intro")
var second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	setup() // Setup the demo (flags etc.)

	// Show intro
	if !*skipIntro {
		introScreen()
		clear()
	}

	showcase("Structured Logging", 5, func() {
		logger := pterm.DefaultLogger.
			WithLevel(pterm.LogLevelTrace)

		logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

		time.Sleep(time.Second * 3)

		interstingStuff := map[string]any{
			"when were crayons invented":  "1903",
			"what is the meaning of life": 42,
			"is this interesting":         true,
		}
		logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
		time.Sleep(time.Second * 3)

		logger.Info("That was actually interesting", logger.Args("such", "wow"))
		time.Sleep(time.Second * 3)
		logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
		time.Sleep(time.Second * 3)
		logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
		time.Sleep(time.Second * 3)
		logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	})

	showcase("Progress bar", 2, func() {
		pb, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
		for i := 0; i < pb.Total; i++ {
			pb.UpdateTitle("Installing " + pseudoProgramList[i])
			if pseudoProgramList[i] == "pseudo-minecraft" {
				pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
			} else {
				pterm.Success.Println("Installing " + pseudoProgramList[i])
			}
			pb.Increment()
			time.Sleep(second / 2)
		}
		pb.Stop()
	})

	showcase("Spinner", 2, func() {
		list := pseudoProgramList[7:]
		spinner, _ := pterm.DefaultSpinner.Start("Installing stuff")
		for i := 0; i < len(list); i++ {
			spinner.UpdateText("Installing " + list[i])
			if list[i] == "pseudo-minecraft" {
				pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
			} else {
				pterm.Success.Println("Installing " + list[i])
			}
			time.Sleep(second)
		}
		spinner.Success()
	})

	showcase("Live Output", 2, func() {
		pterm.Info.Println("You can use an Area to display changing output:")
		pterm.Println()
		area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.
		for i := 0; i < 10; i++ {
			str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender() // Save current time in str.
			area.Update(str)                                                                                              // Update Area contents.
			time.Sleep(time.Second)
		}
		area.Stop()
	})

	showcase("Tables", 4, func() {
		for i := 0; i < 3; i++ {
			pterm.Println()
		}
		td := [][]string{
			{"Library", "Description"},
			{"PTerm", "Make beautiful CLIs"},
			{"Testza", "Programmer friendly test framework"},
			{"Cursor", "Move the cursor around the terminal"},
		}
		table, _ := pterm.DefaultTable.WithHasHeader().WithData(td).Srender()
		boxedTable, _ := pterm.DefaultTable.WithHasHeader().WithData(td).WithBoxed().Srender()
		pterm.DefaultCenter.Println(table)
		pterm.DefaultCenter.Println(boxedTable)
	})

	showcase("TrueColor Support", 7, func() {
		from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
		to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

		str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)\n\nLorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."
		strs := strings.Split(str, "")
		var fadeInfo string // String which will be used to print info.
		// For loop over the range of the string length.
		for i := 0; i < len(str); i++ {
			// Append faded letter to info string.
			fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
		}
		pterm.DefaultCenter.WithCenterEachLineSeparately().Println(fadeInfo)
	})

	showcase("Fully Customizable", 2, func() {
		for i := 0; i < 4; i++ {
			pterm.Println()
		}
		text := "All printers are fully customizable!"
		area := pterm.DefaultArea.WithCenter()
		area.Update(pterm.DefaultBox.Sprintln(text))
		time.Sleep(second)
		area.Update(pterm.DefaultBox.WithTopPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopLeft().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopCenter().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopRight().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomRight().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomCenter().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomLeft().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgCyan)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgRed)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgGreen)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).
			WithBottomPadding(1).
			WithLeftPadding(1).
			WithRightPadding(1).
			WithHorizontalString("═").
			WithVerticalString("║").
			WithBottomLeftCornerString("╗").
			WithBottomRightCornerString("╔").
			WithTopLeftCornerString("╝").
			WithTopRightCornerString("╚").
			Sprintln(text))
		area.Stop()
	})

	showcase("Themes", 2, func() {
		pterm.Info.Println("You can change the color theme of PTerm easily to fit your needs!\nThis is the default one:")
		time.Sleep(second / 2)
		// Print every value of the default theme with its own style.
		v := reflect.ValueOf(pterm.ThemeDefault)
		typeOfS := v.Type()

		if typeOfS == reflect.TypeOf(pterm.Theme{}) {
			for i := 0; i < v.NumField(); i++ {
				field, ok := v.Field(i).Interface().(pterm.Style)
				if ok {
					field.Println(typeOfS.Field(i).Name)
				}
				time.Sleep(time.Millisecond * 250)
			}
		}
	})

	showcase("And much more!", 3, func() {
		for i := 0; i < 4; i++ {
			pterm.Println()
		}
		box := pterm.DefaultBox.
			WithBottomPadding(1).
			WithTopPadding(1).
			WithLeftPadding(3).
			WithRightPadding(3).
			Sprintf("Have fun exploring %s!", pterm.Cyan("PTerm"))
		pterm.DefaultCenter.Println(box)
	})
}

func setup() {
	flag.Parse()
	if *speedup {
		second = time.Millisecond * 200
	}
}

func introScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		putils.LettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("PTDP - PTerm Demo Program"))

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
	time.Sleep(second)
	for i := 14; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(second)
	}
	introSpinner.Stop()
}

func clear() {
	print("\033[H\033[2J")
}

func showcase(title string, seconds int, content func()) {
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithFullWidth().Println(title)
	pterm.Println()
	time.Sleep(second / 2)
	content()
	time.Sleep(second * time.Duration(seconds))
	print("\033[H\033[2J")
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

```

</details>

### header/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a default header.
	// This uses the default settings of PTerm to print a header.
	pterm.DefaultHeader.Println("This is the default header!")

	// Print a spacer line for better readability.
	pterm.Println()

	// Print a full-width header.
	// This uses the WithFullWidth() option of PTerm to print a header that spans the full width of the terminal.
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}

```

</details>

### header/custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/custom/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Customize the DefaultHeader with a cyan background, black text, and a margin of 15.
	pterm.DefaultHeader.WithMargin(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("This is a custom header!")

	// Define a new HeaderPrinter with a red background, black text, and a margin of 20.
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	// Print the custom header using the new HeaderPrinter.
	newHeader.Println("This is a custom header!")
}

```

</details>

### heatmap/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/demo/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/custom_colors/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/custom_legend/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/custom_rgb/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/no_grid/animation.svg)

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

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/heatmap/separated/animation.svg)

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

### interactive_confirm/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_confirm/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Show an interactive confirmation dialog and get the result.
	result, _ := pterm.DefaultInteractiveConfirm.Show()

	// Print a blank line for better readability.
	pterm.Println()

	// Print the user's answer in a formatted way.
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

// boolToText converts a boolean value to a colored text.
// If the value is true, it returns a green "Yes".
// If the value is false, it returns a red "No".
func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

```

</details>

### interactive_continue/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_continue/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Create an interactive continue prompt with default settings
	// This will pause the program execution until the user presses enter
	// The message displayed is "Press 'Enter' to continue..."
	prompt := pterm.DefaultInteractiveContinue

	// Show the prompt and wait for user input
	// The returned result is the user's input (should be empty as it's a continue prompt)
	// The second return value is an error which is ignored here
	result, _ := prompt.Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's input with an info prefix
	// As this is a continue prompt, the input should be empty
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_multiselect/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options.
	var options []string

	// Populate the options slice with 100 options.
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Add 5 more options to the slice, indicating the availability of fuzzy searching.
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Use PTerm's interactive multiselect to present the options to the user and capture their selections.
	// The Show() method displays the options and waits for user input.
	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()

	// Print the selected options, highlighted in green.
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}

```

</details>

### interactive_multiselect/custom-checkmarks

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/custom-checkmarks/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Populate the options slice with 5 options
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Create a new interactive multiselect printer with the options
	// Disable the filter and define the checkmark symbols
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithCheckmark(&pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")})

	// Show the interactive multiselect and get the selected options
	selectedOptions, _ := printer.Show()

	// Print the selected options
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}

```

</details>

### interactive_multiselect/custom-keys

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_multiselect/custom-keys/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"atomicgo.dev/keyboard/keys"
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Populate the options slice with 5 options
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Create a new interactive multiselect printer with the options
	// Disable the filter and set the keys for confirming and selecting options
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithFilter(false).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space)

	// Show the interactive multiselect and get the selected options
	selectedOptions, _ := printer.Show()

	// Print the selected options
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
}

```

</details>

### interactive_select/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_select/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"
	"github.com/pterm/pterm"
)

func main() {
	// Initialize an empty slice to hold the options
	var options []string

	// Generate 100 options and add them to the options slice
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Generate 5 additional options with a specific message and add them to the options slice
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}

```

</details>

### interactive_textinput/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Create an interactive text input with single line input mode
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)

	// Show the text input and get the result
	result, _ := textInput.Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's answer with an info prefix
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_textinput/multi-line

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/multi-line/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Create a default interactive text input with multi-line enabled.
	// This allows the user to input multiple lines of text.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	// Show the text input to the user and store the result.
	// The second return value (an error) is ignored with '_'.
	result, _ := textInput.Show()

	// Print a blank line for better readability in the output.
	pterm.Println()

	// Print the user's input prefixed with an informational message.
	// The '%s' placeholder is replaced with the user's input.
	pterm.Info.Printfln("You answered: %s", result)
}

```

</details>

### interactive_textinput/password

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/interactive_textinput/password/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create an interactive text input with a mask for password input
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")

	// Show the password input prompt and store the result
	result, _ := passwordInput.Show("Enter your password")

	// Get the default logger from PTerm
	logger := pterm.DefaultLogger

	// Log the received password (masked)
	// Note: In a real-world application, you should never log passwords
	logger.Info("Password received", logger.Args("password", result))
}

```

</details>

### logger/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	// Create a logger with trace level
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Log a trace level message
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Pause for 3 seconds
	sleep()

	// Define a map with interesting stuff
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug level message with arguments from the map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	// Pause for 3 seconds
	sleep()

	// Log an info level message
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Pause for 3 seconds
	sleep()

	// Log a warning level message
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Pause for 3 seconds
	sleep()

	// Log an error level message
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Pause for 3 seconds
	sleep()

	// Log an info level message with a long text that will be automatically wrapped
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Pause for 3 seconds
	sleep()

	// Log a fatal level message
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

// Function to pause the execution for 3 seconds
func sleep() {
	time.Sleep(time.Second * 3)
}

```

</details>

### logger/custom-key-styles

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/custom-key-styles/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with a level of Trace or higher.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Define a new style for the "priority" key.
	priorityStyle := map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	}

	// Overwrite all key styles with the new map.
	logger = logger.WithKeyStyles(priorityStyle)

	// Log an info message. The "priority" key will be displayed in red.
	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// Define a new style for the "foo" key.
	fooStyle := *pterm.NewStyle(pterm.FgBlue)

	// Append the new style to the existing ones.
	logger.AppendKeyStyle("foo", fooStyle)

	// Log another info message. The "foo" key will be displayed in blue.
	logger.Info("The foo key should now be blue", logger.Args("priority", "low", "foo", "bar"))
}

```

</details>

### logger/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	// Create a logger with a level of Trace or higher.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Log a trace message with additional arguments.
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug message with arguments from a map.
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	// Log an info message with additional arguments.
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Log a warning message with additional arguments.
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Log an error message with additional arguments.
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Log an info message with additional arguments. PTerm will automatically wrap long logs.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Pause for 2 seconds.
	time.Sleep(time.Second * 2)

	// Log a fatal message with additional arguments. This will terminate the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```

</details>

### logger/json

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/json/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with Trace level and JSON formatter
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithFormatter(pterm.LogFormatterJSON)

	// Log a Trace level message with additional arguments
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff
	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a Debug level message with arguments from the map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	// Log Info, Warn, Error, and Fatal level messages with additional arguments
	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```

</details>

### logger/with-caller

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/with-caller/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with Trace level and caller information
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithCaller()

	// Log a trace message with additional arguments
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff
	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug message with arguments from a map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	// Log an info message with additional arguments
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Log a warning message with additional arguments
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Log an error message with additional arguments
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Log an info message with additional arguments. PTerm will automatically wrap long logs.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Log a fatal message with additional arguments. This will terminate the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```

</details>

### multiple-live-printers/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/multiple-live-printers/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer for managing multiple printers
	multi := pterm.DefaultMultiPrinter

	// Create two spinners with their own writers
	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	// Create five progress bars with their own writers and a total of 100
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	// Start the multi printer
	multi.Start()

	// Increment progress bars and spinners based on certain conditions
	for i := 1; i <= 100; i++ {
		pb1.Increment() // Increment progress bar 1 every iteration

		if i%2 == 0 {
			pb2.Add(3) // Add 3 to progress bar 2 every even iteration
		}

		if i%5 == 0 {
			pb3.Increment() // Increment progress bar 3 every 5th iteration
		}

		if i%10 == 0 {
			pb4.Increment() // Increment progress bar 4 every 10th iteration
		}

		if i%3 == 0 {
			pb5.Increment() // Increment progress bar 5 every 3rd iteration
		}

		if i%50 == 0 {
			spinner1.Success("Spinner 1 is done!") // Mark spinner 1 as successful every 50th iteration
		}

		if i%60 == 0 {
			spinner2.Fail("Spinner 2 failed!") // Mark spinner 2 as failed every 60th iteration
		}

		time.Sleep(time.Millisecond * 50) // Sleep for 50 milliseconds between each iteration
	}

	// Stop the multi printer
	multi.Stop()
}

```

</details>

### panel/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/panel/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define panels in a 2D grid system
	panels := pterm.Panels{
		{
			{Data: "This is the first panel"},
			{Data: pterm.DefaultHeader.Sprint("Hello, World!")},
			{Data: "This\npanel\ncontains\nmultiple\nlines"},
		},
		{
			{Data: pterm.Red("This is another\npanel line")},
			{Data: "This is the second panel\nwith a new line"},
		},
	}

	// Render the panels with a padding of 5
	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}

```

</details>

### paragraph/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/paragraph/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Using the default paragraph printer to print a long text.
	// The text is split at the spaces, which is useful for continuous text of all kinds.
	// The line width can be manually adjusted if needed.
	pterm.DefaultParagraph.Println("This is the default paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")

	// Printing a line space for separation.
	pterm.Println()

	// Printing a long text without using the paragraph printer.
	// The default Println() function is used here, which does not provide intelligent splitting.
	pterm.Println("This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")
}

```

</details>

### paragraph/customized

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/paragraph/customized/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a long text to be printed as a paragraph.
	longText := "This is a custom paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// Print the long text as a paragraph with a custom maximal width of 60 characters.
	pterm.DefaultParagraph.WithMaxWidth(60).Println(longText)

	// Print a line space to separate the paragraph from the following text.
	pterm.Println()

	// Define another long text to be printed without a paragraph printer.
	longTextWithoutParagraph := "This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// Print the long text without using a paragraph printer.
	pterm.Println(longTextWithoutParagraph)
}

```

</details>

### prefix/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/prefix/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Enable debug messages in PTerm.
	pterm.EnableDebugMessages()

	// Print a debug message with PTerm.
	pterm.Debug.Println("Hello, World!")

	// Print an informational message with PTerm.
	pterm.Info.Println("Hello, World!")

	// Print a success message with PTerm.
	pterm.Success.Println("Hello, World!")

	// Print a warning message with PTerm.
	pterm.Warning.Println("Hello, World!")

	// Print an error message with PTerm. This will also display the filename and line number in the terminal.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!")

	// Print an informational message with PTerm, with line number.
	// This demonstrates that other PrefixPrinters can also display line numbers.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")

	// Temporarily set Fatal to false, so that the CI won't crash.
	// This will print a fatal message with PTerm, but won't terminate the program.
	pterm.Fatal.WithFatal(false).Println("Hello, World!")
}

```

</details>

### progressbar/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Slice of strings representing names of pseudo applications to be downloaded.
var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	// Create a progressbar with the total steps equal to the number of items in fakeInstallList.
	// Set the initial title of the progressbar to "Downloading stuff".
	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	// Loop over each item in the fakeInstallList.
	for i := 0; i < p.Total; i++ {
		// Simulate a slow download for the 7th item.
		if i == 6 {
			time.Sleep(time.Second * 3)
		}

		// Update the title of the progressbar with the current item being downloaded.
		p.UpdateTitle("Downloading " + fakeInstallList[i])

		// Print a success message for the current download. This will be printed above the progressbar.
		pterm.Success.Println("Downloading " + fakeInstallList[i])

		// Increment the progressbar by one to indicate progress.
		p.Increment()

		// Pause for 350 milliseconds to simulate the time taken for each download.
		time.Sleep(time.Millisecond * 350)
	}
}

```

</details>

### progressbar/multiple

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/multiple/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer instance from the default one
	multi := pterm.DefaultMultiPrinter

	// Create five progress bars with a total of 100 units each, and assign each a new writer from the multi printer
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	// Start the multi printer
	multi.Start()

	// Loop to increment progress bars based on certain conditions
	for i := 1; i <= 100; i++ {
		pb1.Increment() // Increment the first progress bar at each iteration

		if i%2 == 0 {
			pb2.Add(3) // Add 3 units to the second progress bar at every even iteration
		}

		if i%5 == 0 {
			pb3.Increment() // Increment the third progress bar at every fifth iteration
		}

		if i%10 == 0 {
			pb4.Increment() // Increment the fourth progress bar at every tenth iteration
		}

		if i%3 == 0 {
			pb5.Increment() // Increment the fifth progress bar at every third iteration
		}

		time.Sleep(time.Millisecond * 50) // Pause for 50 milliseconds at each iteration
	}

	// Stop the multi printer
	multi.Stop()
}

```

</details>

### section/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/section/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a section with level one and print it.
	pterm.DefaultSection.Println("This is a section!")

	// Print an informational message.
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")

	// Create a section with level two and print it.
	pterm.DefaultSection.WithLevel(2).Println("This is another section!")

	// Print another informational message.
	pterm.Info.Println("And this is\nmore placeholder text")
}

```

</details>

### slog/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/slog/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"log/slog"

	"github.com/pterm/pterm"
)

func main() {
	// Create a new slog handler with the default PTerm logger
	handler := pterm.NewSlogHandler(&pterm.DefaultLogger)

	// Create a new slog logger with the handler
	logger := slog.New(handler)

	// Log a debug message (won't show by default)
	logger.Debug("This is a debug message that won't show")

	// Change the log level to debug to enable debug messages
	pterm.DefaultLogger.Level = pterm.LogLevelDebug

	// Log a debug message (will show because debug level is enabled)
	logger.Debug("This is a debug message", "changedLevel", true)

	// Log an info message
	logger.Info("This is an info message")

	// Log a warning message
	logger.Warn("This is a warning message")

	// Log an error message
	logger.Error("This is an error message")
}

```

</details>

### spinner/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create and start a fork of the default spinner.
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Some informational action...")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerInfo.Info()          // Resolve spinner with error message.

	// Create and start a fork of the default spinner.
	spinnerSuccess, _ := pterm.DefaultSpinner.Start("Doing something important... (will succeed)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerSuccess.Success()    // Resolve spinner with success message.

	// Create and start a fork of the default spinner.
	spinnerWarning, _ := pterm.DefaultSpinner.Start("Doing something important... (will warn)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerWarning.Warning()    // Resolve spinner with warning message.

	// Create and start a fork of the default spinner.
	spinnerFail, _ := pterm.DefaultSpinner.Start("Doing something important... (will fail)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerFail.Fail()          // Resolve spinner with error message.

	// Create and start a fork of the default spinner.
	spinnerNochange, _ := pterm.DefaultSpinner.Start("Checking something important... (will result in no change)")
	// Replace the InfoPrinter with a custom "NOCHG" one
	spinnerNochange.InfoPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.Style{pterm.FgLightBlue},
		Prefix: pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgLightBlue},
			Text:  " NOCHG ",
		},
	}
	time.Sleep(time.Second * 2)                     // Simulate 3 seconds of processing something.
	spinnerNochange.Info("No change were required") // Resolve spinner with error message.

	// Create and start a fork of the default spinner.
	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Doing a lot of stuff...")
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.UpdateText("It's really much")   // Update spinner text.
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.UpdateText("We're nearly done!") // Update spinner text.
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.Success("Finally!")              // Resolve spinner with success message.
}

```

</details>

### spinner/multiple

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/multiple/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer. This allows multiple spinners to print simultaneously.
	multi := pterm.DefaultMultiPrinter

	// Create and start spinner 1 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 1".
	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")

	// Create and start spinner 2 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 2".
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	// Create and start spinner 3 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 3".
	spinner3, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 3")

	// Start the multi printer. This will start printing all the spinners.
	multi.Start()

	// Wait for 1 second.
	time.Sleep(time.Millisecond * 1000)

	// Stop spinner 1 with a success message.
	spinner1.Success("Spinner 1 is done!")

	// Wait for 750 milliseconds.
	time.Sleep(time.Millisecond * 750)

	// Stop spinner 2 with a failure message.
	spinner2.Fail("Spinner 2 failed!")

	// Wait for 500 milliseconds.
	time.Sleep(time.Millisecond * 500)

	// Stop spinner 3 with a warning message.
	spinner3.Warning("Spinner 3 has a warning!")

	// Stop the multi printer. This will stop printing all the spinners.
	multi.Stop()
}

```

</details>

### style/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/style/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a primary style with light cyan foreground, gray background, and bold text
	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)

	// Define a secondary style with light green foreground, white background, and italic text
	secondary := pterm.NewStyle(pterm.FgLightGreen, pterm.BgWhite, pterm.Italic)

	// Print "Hello, World!" with the primary style
	primary.Println("Hello, World!")

	// Print "Hello, World!" with the secondary style
	secondary.Println("Hello, World!")
}

```

</details>

### table/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the first table
	tableData1 := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

	pterm.Println() // Blank line

	// Define the data for the second table
	tableData2 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Create another table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithData(tableData2).Render()
}

```

</details>

### table/boxed

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/boxed/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	// Each inner slice represents a row in the table.
	// The first row is considered as the header of the table.
	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with the defined data.
	// The table has a header and is boxed.
	// Finally, render the table to print it.
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
}

```

</details>

### table/multiple-lines

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/multiple-lines/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	data := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Create and render the table.
	// The options are chained in a single line for simplicity.
	// The table has a header, a row separator, and a header row separator.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(data).Render()
}

```

</details>

### table/right-alignment

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/right-alignment/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	// Each inner slice represents a row in the table.
	// The first row is considered as the header.
	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with the defined data.
	// The table has a header and the text in the cells is right-aligned.
	// The Render() method is used to print the table to the console.
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(tableData).Render()
}

```

</details>

### theme/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/theme/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"reflect"
	"time"
)

func main() {
	// Print an informational message about the default theme styles.
	pterm.Info.Println("These are the default theme styles.\nYou can modify them easily to your personal preference,\nor create new themes from scratch :)")

	// Print a blank line for better readability.
	pterm.Println()

	// Get the value and type of the default theme.
	v := reflect.ValueOf(pterm.ThemeDefault)
	typeOfS := v.Type()

	// Check if the type of the default theme is 'pterm.Theme'.
	if typeOfS == reflect.TypeOf(pterm.Theme{}) {
		// Iterate over each field in the default theme.
		for i := 0; i < v.NumField(); i++ {
			// Try to convert the field to 'pterm.Style'.
			field, ok := v.Field(i).Interface().(pterm.Style)
			if ok {
				// Print the field name using its own style.
				field.Println(typeOfS.Field(i).Name)
			}
			// Pause for a quarter of a second to make the output easier to read.
			time.Sleep(time.Millisecond * 250)
		}
	}
}

```

</details>

### tree/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/tree/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define a tree structure using pterm.TreeNode
	tree := pterm.TreeNode{
		// The top node of the tree
		Text: "Top node",
		// The children of the top node
		Children: []pterm.TreeNode{{
			// A child node
			Text: "Child node",
			// The children of the child node
			Children: []pterm.TreeNode{
				// Grandchildren nodes
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
			},
		}},
	}

	// Render the tree with the defined structure as the root
	pterm.DefaultTree.WithRoot(tree).Render()
}

```

</details>

### tree/from-leveled-list

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/tree/from-leveled-list/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Define a leveled list to represent the structure of the directories.
	leveledList := pterm.LeveledList{
		{Level: 0, Text: "C:"},
		{Level: 1, Text: "Users"},
		{Level: 1, Text: "Windows"},
		{Level: 1, Text: "Programs"},
		{Level: 1, Text: "Programs(x86)"},
		{Level: 1, Text: "dev"},
		{Level: 0, Text: "D:"},
		{Level: 0, Text: "E:"},
		{Level: 1, Text: "Movies"},
		{Level: 1, Text: "Music"},
		{Level: 2, Text: "LinkinPark"},
		{Level: 1, Text: "Games"},
		{Level: 2, Text: "Shooter"},
		{Level: 3, Text: "CallOfDuty"},
		{Level: 3, Text: "CS:GO"},
		{Level: 3, Text: "Battlefield"},
		{Level: 4, Text: "Battlefield 1"},
		{Level: 4, Text: "Battlefield 2"},
		{Level: 0, Text: "F:"},
		{Level: 1, Text: "dev"},
		{Level: 2, Text: "dops"},
		{Level: 2, Text: "PTerm"},
	}

	// Convert the leveled list into a tree structure.
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = "Computer" // Set the root node text.

	// Render the tree structure using the default tree printer.
	pterm.DefaultTree.WithRoot(root).Render()
}

```

</details>


<!-- examples:end -->


---

> GitHub [@pterm](https://github.com/pterm) &nbsp;&middot;&nbsp;
> Author [@MarvinJWendt](https://github.com/MarvinJWendt)
> | [PTerm.sh](https://pterm.sh)




































































































































