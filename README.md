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

If you want to create a CLI tool, make sure to check out our [cli-template](https://github.com/pterm/cli-template),
which features automatic website generation, automatic deployments, a custom CI-System and much more!

## ⭐ Main Features

|Feature|Description|
|-------|-----------|
|🪀 Easy to use |Our first priority is to keep PTerm as easy to use as possible.<br> With many [examples](#-examples) for each individual component, getting started with PTerm is extremely easy.<br> All components are similar in design and implement interfaces to simplify mixing individual components together.|
|🤹‍♀️ Cross-Platform |We take special precautions to ensure that PTerm works on as many operating systems and terminals as possible.<br> Whether it's `Windows CMD`, `macOS iTerm2` or in the backend (for example inside a `GitHub Action` or other CI systems), PTerm **guarantees** beautiful output!|
|🧪 Well tested |PTerm has a 100% test coverage, which means that every line of code inside PTerm gets tested automatically<br>We test PTerm continuously. However, since a human cannot test everything all the time, we have our own test system with which we currently run <!-- unittestcount2:start -->**`28774`**<!-- unittestcount2:end -->automated tests to ensure that PTerm has no bugs. |
|✨ Consistent Colors|PTerm uses the [ANSI color scheme](https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit) which is widely used by terminals to ensure consistent colors in different terminal themes.<br>If that's not enough, PTerm can be used to access the full RGB color scheme (16 million colors) in terminals that support `TrueColor`.|
|📚 Component system|PTerm consists of many components, called `Printers`, which can be used individually or together to generate pretty console output.|
|🛠 Configurable|PTerm can be used by without any configuration. However, you can easily configure each component with little code, so everyone has the freedom to design their own terminal output.|
|✏ Documentation |To view the official documentation of the latest release, you can go to the automatically generated page of [pkg.go.dev](https://pkg.go.dev/github.com/pterm/pterm#section-documentation) This documentation is very technical and includes every method that can be used in PTerm.<br>**For an easy start we recommend that you take a look at the [examples section](#-examples).** Here you can see pretty much every feature of PTerm with example code. The animations of the examples are automatically updated as soon as something changes in PTerm.|

<div align="center">

### Printers (Components)

|Feature|Examples| - |Feature|Examples|
|-------|--------|---|-----|--------|
|Bar Charts|[Examples](https://github.com/pterm/pterm/tree/master/_examples/barchart)|-|RGB|[Examples](https://github.com/pterm/pterm/tree/master/_examples/coloring)|
|BigText|[Examples](https://github.com/pterm/pterm/tree/master/_examples/bigtext)|-|Sections|[Examples](https://github.com/pterm/pterm/tree/master/_examples/section)|
|Box|[Examples](https://github.com/pterm/pterm/tree/master/_examples/box)|-|Spinners|[Examples](https://github.com/pterm/pterm/tree/master/_examples/spinner)|
|Bullet Lists|[Examples](https://github.com/pterm/pterm/tree/master/_examples/bulletlist)|-|Trees|[Examples](https://github.com/pterm/pterm/tree/master/_examples/tree)|
|Centered|[Examples](https://github.com/pterm/pterm/tree/master/_examples/center)|-|Theming|[Examples](https://github.com/pterm/pterm/tree/master/_examples/theme)|
|Colors|[Examples](https://github.com/pterm/pterm/tree/master/_examples/coloring)|-|Tables|[Examples](https://github.com/pterm/pterm/tree/master/_examples/table)|
|Headers|[Examples](https://github.com/pterm/pterm/tree/master/_examples/header)|-|Styles|[Examples](https://github.com/pterm/pterm/tree/master/_examples/style)|
|Panels|[Examples](https://github.com/pterm/pterm/tree/master/_examples/panel)|-|Area|[Examples](https://github.com/pterm/pterm/tree/master/_examples/area)|
|Paragraphs|[Examples](https://github.com/pterm/pterm/tree/master/_examples/paragraph)|-|||
|Prefixes|[Examples](https://github.com/pterm/pterm/tree/master/_examples/prefix)|-|||
|Progress Bars|[Examples](https://github.com/pterm/pterm/tree/master/_examples/progressbar)|-|||

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
		putils.LettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	// LettersFromStringWithRGB can be used to create a large text with a specific RGB color.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0))).
		Render()
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
	pterm.NewRGB(201, 144, 30).Println("This text is printed with a custom RGB!")
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
//   go run main.go -speedup
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

	showcase("Default Prefix Printers", 5, func() {
		// Enable debug messages.
		pterm.EnableDebugMessages() // Temporarily set debug output to true, to display the debug printer.

		pterm.Debug.Println("Hello, World!") // Print Debug.
		time.Sleep(second / 2)
		pterm.Info.Println("Hello, World!") // Print Info.
		time.Sleep(second / 2)
		pterm.Success.Println("Hello, World!") // Print Success.
		time.Sleep(second / 2)
		pterm.Warning.Println("Hello, World!") // Print Warning.
		time.Sleep(second / 2)
		pterm.Error.Println("Errors show the filename and linenumber inside the terminal!") // Print Error.
		time.Sleep(second / 2)
		pterm.Info.WithShowLineNumber().Println("Other PrefixPrinters can do that too!") // Print Error.
		time.Sleep(second / 2)
		// Temporarily set Fatal to false, so that the CI won't panic.
		pterm.Fatal.WithFatal(false).Println("Hello, World!") // Print Fatal.

		pterm.DisableDebugMessages() // Disable debug output again.
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
	pterm.DefaultHeader.Println("This is the default header!")
	pterm.Println() // spacer
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}

```

</details>

### header-custom/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header-custom/demo/animation.svg)

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
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com", "just a test, hey!"},
	}).Render()

	pterm.Println() // Blank line

	// Create a table with right alignment.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}).WithRightAlignment().Render()
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




































































































































