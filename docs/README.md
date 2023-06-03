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
<p align="center">Show Demo Code</p>
</a>

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

<!-- printers:start -->
| Feature | Feature | Feature | Feature | Feature |
| :-------: | :-------: | :-------: | :-------: | :-------: |
| Area <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/area) |Barchart <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/barchart) |Basictext <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/basictext) |Bigtext <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/bigtext) |Box <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/box) |
| Bulletlist <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/bulletlist) |Center <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/center) |Coloring <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/coloring) |Demo <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/demo) |Header <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/header) |
| Interactive confirm <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_confirm) |Interactive continue <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_continue) |Interactive multiselect <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_multiselect) |Interactive select <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_select) |Interactive textinput <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/interactive_textinput) |
| Logger <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/logger) |Panel <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/panel) |Paragraph <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/paragraph) |Prefix <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/prefix) |Progressbar <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/progressbar) |
| Section <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/section) |Spinner <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/spinner) |Style <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/style) |Table <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/table) |Theme <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/theme) |
| Tree <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/tree) | |  |  |  | 
<!-- printers:end -->


<div align="center">

### 🦸‍♂️ Supporters

|-|User|💸|
|---|---|---|
|![Jens Lauterbach](https://avatars.githubusercontent.com/u/1292368?s=25)|[@jenslauterbach](https://github.com/jenslauterbach)|25$|

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

### basictext/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/basictext/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A BasicText printer is used to print text, without special formatting.
	// As it implements the TextPrinter interface, you can use it in combination with other printers.
	pterm.DefaultBasicText.Println("Default basic text printer.")
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
	pterm.DefaultBasicText.Println("For example to resolve progressbars and spinners.")
	// If you just want to print text, you should use this instead:
	// 	pterm.Println("Hello, World!")
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
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()
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
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()
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
	// Print a large text with the LetterStyle from the standard theme.
	// Useful for title screens.
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()

	// Print a large text with differently colored letters.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()

	// LettersFromStringWithRGB can be used to create a large text with a specific RGB color.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0))).
		Render()
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
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println("Hello, World!")
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
	pterm.DefaultBox.Println("Hello, World!")
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
	pterm.Info.Println("This might not be rendered correctly on GitHub,\nbut it will work in a real terminal.\nThis is because GitHub does not use a monospaced font by default for SVGs")

	panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	panel3 := pterm.DefaultBox.WithTitle("bottom center title").WithTitleBottomCenter().Sprint("Duis aute irure\ndolor in reprehenderit\nin voluptate velit esse cillum\ndolore eu fugiat\nnulla pariatur.")

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: panel1}, {Data: panel2}},
		{{Data: panel3}},
	}).Srender()

	pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
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
	// Default titled bpx
	paddedBox := pterm.DefaultBox.WithLeftPadding(4).WithRightPadding(4).WithTopPadding(1).WithBottomPadding(1)

	title := pterm.LightRed("I'm a box!")

	box1 := paddedBox.WithTitle(title).Sprint("Hello, World!\n      1")
	box2 := paddedBox.WithTitle(title).WithTitleTopCenter().Sprint("Hello, World!\n      2")
	box3 := paddedBox.WithTitle(title).WithTitleTopRight().Sprint("Hello, World!\n      3")
	box4 := paddedBox.WithTitle(title).WithTitleBottomRight().Sprint("Hello, World!\n      4")
	box5 := paddedBox.WithTitle(title).WithTitleBottomCenter().Sprint("Hello, World!\n      5")
	box6 := paddedBox.WithTitle(title).WithTitleBottomLeft().Sprint("Hello, World!\n      6")
	box7 := paddedBox.WithTitle(title).WithTitleTopLeft().Sprint("Hello, World!\n      7")

	pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{box1}, {box2}, {box3}},
		{{box4}, {box5}, {box6}},
		{{box7}},
	}).Render()
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
	// Print a customized list with different styles and levels.
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Blue", TextStyle: pterm.NewStyle(pterm.FgBlue), BulletStyle: pterm.NewStyle(pterm.FgRed)},
		{Level: 1, Text: "Green", TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		{Level: 2, Text: "Cyan", TextStyle: pterm.NewStyle(pterm.FgCyan), Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgYellow)},
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
	// Print a list with different levels.
	// Useful to generate lists automatically from data.
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}).Render()

	// Convert a text to a list and print it.
	putils.BulletListFromString(`0
 1
  2
   3`, " ").Render()
}

```

</details>

### center/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/center/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultCenter.Println("This text is centered!\nIt centeres the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("PTerm")).Srender()
	pterm.DefaultCenter.Println(s) // Print BigLetters with the default CenterPrinter

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
	// Print all colors

	pterm.DefaultTable.WithData([][]string{
		{pterm.FgBlack.Sprint("Black"), pterm.FgRed.Sprint("Red"), pterm.FgGreen.Sprint("Green"), pterm.FgYellow.Sprint("Yellow")},
		{"", pterm.FgLightRed.Sprint("Light Red"), pterm.FgLightGreen.Sprint("Light Green"), pterm.FgLightYellow.Sprint("Light Yellow")},
		{pterm.BgBlack.Sprint("Black"), pterm.BgRed.Sprint("Red"), pterm.BgGreen.Sprint("Green"), pterm.BgYellow.Sprint("Yellow")},
		{"", pterm.BgLightRed.Sprint("Light Red"), pterm.BgLightGreen.Sprint("Light Green"), pterm.BgLightYellow.Sprint("Light Yellow")},
		{pterm.FgBlue.Sprint("Blue"), pterm.FgMagenta.Sprint("Magenta"), pterm.FgCyan.Sprint("Cyan"), pterm.FgWhite.Sprint("White")},
		{pterm.FgLightBlue.Sprint("Light Blue"), pterm.FgLightMagenta.Sprint("Light Magenta"), pterm.FgLightCyan.Sprint("Light Cyan"), pterm.FgLightWhite.Sprint("Light White")},
		{pterm.BgBlue.Sprint("Blue"), pterm.BgMagenta.Sprint("Magenta"), pterm.BgCyan.Sprint("Cyan"), pterm.BgWhite.Sprint("White")},
		{pterm.BgLightBlue.Sprint("Light Blue"), pterm.BgLightMagenta.Sprint("Light Magenta"), pterm.BgLightCyan.Sprint("Light Cyan"), pterm.BgLightWhite.Sprint("Light White")},
	}).Render()

	pterm.Println()

	// Print different colored words.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))

	pterm.Println()

	// Or print colors as a style
	style := pterm.NewStyle(pterm.BgRed, pterm.FgLightGreen, pterm.Bold)
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
	for i := 0; i < 15; i++ {
		switch i {
		case 5:
			pterm.Info.Println("Disabled Output!")
			pterm.DisableOutput()
		case 10:
			pterm.EnableOutput()
			pterm.Info.Println("Enabled Output!")
		}

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
	// Print info.
	pterm.Info.Println("RGB colors only work in Terminals which support TrueColor.")

	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients end point.

	// For loop over the range of the terminal height.
	for i := 0; i < pterm.GetTerminalHeight()-2; i++ {
		// Print string which is colored with the faded RGB value.
		from.Fade(0, float32(pterm.GetTerminalHeight()-2), float32(i), to).Println("Hello, World!")
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
	white := pterm.NewRGB(255, 255, 255) // This RGB value is used as the gradients start point.
	grey := pterm.NewRGB(128, 128, 128)  // This RGB value is used as the gradients start point.
	black := pterm.NewRGB(0, 0, 0)       // This RGB value is used as the gradients start point.
	red := pterm.NewRGB(255, 0, 0)       // This RGB value is used as the gradients start point.
	purple := pterm.NewRGB(255, 0, 255)  // This RGB value is used as the gradients start point.
	green := pterm.NewRGB(0, 255, 0)     // This RGB value is used as the gradients start point.

	str := "RGB colors only work in Terminals which support TrueColor."
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += pterm.NewRGBStyle(white.Fade(0, float32(len(str)), float32(i), purple), grey.Fade(0, float32(len(str)), float32(i), black)).Sprint(strs[i])
	}

	pterm.Info.Println(fadeInfo)

	str = "The background and foreground colors can be customized individually."
	strs = strings.Split(str, "")
	var fade2 string // String which will be used to print info.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fade2 += pterm.NewRGBStyle(black, purple.Fade(0, float32(len(str)), float32(i), red)).Sprint(strs[i])
	}

	pterm.Println(fade2)

	str = "Styles can also be applied. For example: Bold or Italic."
	strs = strings.Split(str, "")
	var fade3 string // String which will be used to print.

	bold := 0
	boldStr := strings.Split("Bold", "")
	italic := 0
	italicStr := strings.Split("Italic", "")

	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		s := pterm.NewRGBStyle(white.Fade(0, float32(len(str)), float32(i), green), red.Fade(0, float32(len(str)), float32(i), black))

		// if the next letters are "Bold", then add the style "Bold".
		// else if the next letters are "Italic", then add the style "Italic".
		if bold < len(boldStr) && i+len(boldStr) <= len(strs) {
			if strings.Join(strs[i:i+len(boldStr)-bold], "") == strings.Join(boldStr[bold:], "") {
				s = s.AddOptions(pterm.Bold)
				bold++
			}
		} else if italic < len(italicStr) && i+len(italicStr)-italic < len(strs) {
			if strings.Join(strs[i:i+len(italicStr)-italic], "") == strings.Join(italicStr[italic:], "") {
				s = s.AddOptions(pterm.Italic)
				italic++
			}
		}
		fade3 += s.Sprint(strs[i])
	}

	pterm.Println(fade3)
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
	from := pterm.NewRGB(0, 255, 255)  // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)    // This RGB value is used as the gradients first point.
	to2 := pterm.NewRGB(255, 0, 0)     // This RGB value is used as the gradients second point.
	to3 := pterm.NewRGB(0, 255, 0)     // This RGB value is used as the gradients third point.
	to4 := pterm.NewRGB(255, 255, 255) // This RGB value is used as the gradients end point.

	str := "RGB colors only work in Terminals which support TrueColor."
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}

	// Print info.
	pterm.Info.Println(fadeInfo)

	// For loop over the range of the terminal height.
	for i := 0; i < pterm.GetTerminalHeight()-2; i++ {
		// Print string which is colored with the faded RGB value.
		from.Fade(0, float32(pterm.GetTerminalHeight()-2), float32(i), to, to2, to3, to4).Println("Hello, World!")
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
	// Print default error.
	pterm.Error.Println("This is the default Error")

	// Customize default error.
	pterm.Error.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed),
	}

	// Print new default error.
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
	// Print strings with a custom RGB color.
	// NOTICE: This only works with terminals which support TrueColor.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")
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
	foregroundRGB := pterm.RGB{R: 187, G: 80, B: 0}
	backgroundRGB := pterm.RGB{R: 0, G: 50, B: 123}

	// Print string with a custom foreground and background RGB color.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).Println("This text is not styled.")

	// Print string with a custom foreground and background RGB color and style bold.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).AddOptions(pterm.Bold).Println("This text is bold.")

	// Print string with a custom foreground and background RGB color and style italic.
	pterm.NewRGBStyle(foregroundRGB, backgroundRGB).AddOptions(pterm.Italic).Println("This text is italic.")
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

	showcase("Fully Customizale", 2, func() {
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

### header/custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/custom/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// All available options: https://pkg.go.dev/github.com/pterm/pterm#HeaderPrinter

	// Build on top of DefaultHeader
	pterm.DefaultHeader. // Use DefaultHeader as base
				WithMargin(15).
				WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).
				WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
				Println("This is a custom header!")
	// Instead of printing the header you can set it to a variable.
	// You can then reuse your custom header.

	// Making a completely new HeaderPrinter
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	// Print header.
	newHeader.Println("This is a custom header!")
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
	pterm.DefaultHeader.Println("This is the default header!")
	pterm.Println() // spacer
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
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
	result, _ := pterm.DefaultInteractiveConfirm.Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

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
	result, _ := pterm.DefaultInteractiveContinue.Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)
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

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

func main() {
	var options []string

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	printer.Checkmark = &pterm.Checkmark{Checked: pterm.Green("+"), Unchecked: pterm.Red("-")}
	selectedOptions, _ := printer.Show()
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
	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func main() {
	var options []string

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	printer := pterm.DefaultInteractiveMultiselect.WithOptions(options)
	printer.Filter = false
	printer.KeyConfirm = keys.Enter
	printer.KeySelect = keys.Space
	selectedOptions, _ := printer.Show()
	pterm.Info.Printfln("Selected options: %s", pterm.Green(selectedOptions))
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
	var options []string

	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	selectedOptions, _ := pterm.DefaultInteractiveMultiselect.WithOptions(options).Show()
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
	var options []string

	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
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
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show()
	pterm.Println() // Blank line
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
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show() // Text input with multi line enabled
	pterm.Println()                                                       // Blank line
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
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Enter your password")

	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
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
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace) // Only show logs with a level of Trace or higher.

	// Overwrite all key styles with a new map
	logger = logger.WithKeyStyles(map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	})

	// The priority key should now be red
	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// Append a key style to the exisiting ones
	logger.AppendKeyStyle("foo", *pterm.NewStyle(pterm.FgBlue))

	// The foo key should now be blue
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
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace) // Only show logs with a level of Trace or higher.

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// You can also use the `ArgsFromMap` function to create a `Args` object from a map.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	time.Sleep(time.Second * 2)
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace)

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	sleep()

	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
	sleep()

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	sleep()
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	sleep()
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	sleep()
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	sleep()
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

func sleep() {
	time.Sleep(time.Second * 3)
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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace).       // Only show logs with a level of Trace or higher.
		WithFormatter(pterm.LogFormatterJSON) // ! Make the logger print JSON logs.

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// You can also use the `ArgsFromMap` function to create a `Args` object from a map.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace). // Only show logs with a level of Trace or higher.
		WithCaller()                    // ! Show the caller of the log function.

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// You can also use the `ArgsFromMap` function to create a `Args` object from a map.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
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
	// Declare panels in a two dimensional grid system.
	panels := pterm.Panels{
		{{Data: "This is the first panel"}, {Data: pterm.DefaultHeader.Sprint("Hello, World!")}, {Data: "This\npanel\ncontains\nmultiple\nlines"}},
		{{Data: pterm.Red("This is another\npanel line")}, {Data: "This is the second panel\nwith a new line"}},
	}

	// Print panels.
	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
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
	// Print a paragraph with a custom maximal width.
	pterm.DefaultParagraph.WithMaxWidth(60).Println("This is a custom paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")

	// Print one line space.
	pterm.Println()

	// Print text without a paragraph printer.
	pterm.Println("This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")
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
	// Print long text with default paragraph printer.
	pterm.DefaultParagraph.Println("This is the default paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")

	// Print one line space.
	pterm.Println()

	// Print long text without paragraph printer.
	pterm.Println("This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")
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
	// Enable debug messages.
	pterm.EnableDebugMessages()

	pterm.Debug.Println("Hello, World!")                                                // Print Debug.
	pterm.Info.Println("Hello, World!")                                                 // Print Info.
	pterm.Success.Println("Hello, World!")                                              // Print Success.
	pterm.Warning.Println("Hello, World!")                                              // Print Warning.
	pterm.Error.Println("Errors show the filename and linenumber inside the terminal!") // Print Error.
	pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!")    // Print Error.
	// Temporarily set Fatal to false, so that the CI won't crash.
	pterm.Fatal.WithFatal(false).Println("Hello, World!") // Print Fatal.
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

// Slice of strings with placeholder text.
var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	// Create progressbar as fork from the default progressbar.
	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		if i == 6 {
			time.Sleep(time.Second * 3) // Simulate a slow download.
		}
		p.UpdateTitle("Downloading " + fakeInstallList[i])         // Update the title of the progressbar.
		pterm.Success.Println("Downloading " + fakeInstallList[i]) // If a progressbar is running, each print will be printed above the progressbar.
		p.Increment()                                              // Increment the progressbar by one. Use Add(x int) to increment by a custom amount.
		time.Sleep(time.Millisecond * 350)                         // Sleep 350 milliseconds.
	}
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
	// Print a section with level one.
	pterm.DefaultSection.Println("This is a section!")
	// Print placeholder.
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")

	// Print a section with level two.
	pterm.DefaultSection.WithLevel(2).Println("This is another section!")
	// Print placeholder.
	pterm.Info.Println("And this is\nmore placeholder text")
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

### style/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/style/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create styles as new variables
	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)
	secondary := pterm.NewStyle(pterm.FgLightGreen, pterm.BgWhite, pterm.Italic)

	// Use created styles
	primary.Println("Hello, World!")
	secondary.Println("Hello, World!")
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
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}).Render()
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
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}).Render()

	pterm.Println() // Blank line

	// Create a table with multiple lines in a row.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}).Render()
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
	// Create a table with multiple lines in a row and set a row separator.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}).Render()
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
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}).Render()
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
	// Print info.
	pterm.Info.Println("These are the default theme styles.\n" +
		"You can modify them easily to your personal preference,\n" +
		"or create new themes from scratch :)")

	pterm.Println() // Print one line space.

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
	tree := pterm.TreeNode{
		Text: "Top node",
		Children: []pterm.TreeNode{{
			Text: "Child node",
			Children: []pterm.TreeNode{
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
				{Text: "Grandchild node"},
			},
		}},
	}

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
	// You can use a LeveledList here, for easy generation.
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Users"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
		pterm.LeveledListItem{Level: 1, Text: "Programs(x86)"},
		pterm.LeveledListItem{Level: 1, Text: "dev"},
		pterm.LeveledListItem{Level: 0, Text: "D:"},
		pterm.LeveledListItem{Level: 0, Text: "E:"},
		pterm.LeveledListItem{Level: 1, Text: "Movies"},
		pterm.LeveledListItem{Level: 1, Text: "Music"},
		pterm.LeveledListItem{Level: 2, Text: "LinkinPark"},
		pterm.LeveledListItem{Level: 1, Text: "Games"},
		pterm.LeveledListItem{Level: 2, Text: "Shooter"},
		pterm.LeveledListItem{Level: 3, Text: "CallOfDuty"},
		pterm.LeveledListItem{Level: 3, Text: "CS:GO"},
		pterm.LeveledListItem{Level: 3, Text: "Battlefield"},
		pterm.LeveledListItem{Level: 4, Text: "Battlefield 1"},
		pterm.LeveledListItem{Level: 4, Text: "Battlefield 2"},
		pterm.LeveledListItem{Level: 0, Text: "F:"},
		pterm.LeveledListItem{Level: 1, Text: "dev"},
		pterm.LeveledListItem{Level: 2, Text: "dops"},
		pterm.LeveledListItem{Level: 2, Text: "PTerm"},
	}

	// Generate tree from LeveledList.
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = "Computer"

	// Render TreePrinter
	pterm.DefaultTree.WithRoot(root).Render()
}

```

</details>


<!-- examples:end -->


---

> GitHub [@pterm](https://github.com/pterm) &nbsp;&middot;&nbsp;
> Author [@MarvinJWendt](https://github.com/MarvinJWendt)
> | [PTerm.sh](https://pterm.sh)




































































































































