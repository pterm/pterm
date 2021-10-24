<!--suppress HtmlDeprecatedAttribute -->

<h1 align="center">💻 PTerm | Pretty Terminal Printer</h1>
<p align="center">A golang module to print pretty text</p>

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

<a href="https://github.com/pterm/pterm/issues" style="text-decoration: none">
<img src="https://img.shields.io/github/issues/pterm/pterm.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT" style="text-decoration: none">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a href="https://github.com/dops-cli/dops/releases" style="text-decoration: none">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

<a href="https://codecov.io/gh/pterm/pterm" style="text-decoration: none">
<img src="https://img.shields.io/codecov/c/gh/pterm/pterm?color=magenta&logo=codecov&style=flat-square" alt="Downloads">
</a>

<a href="https://codecov.io/gh/pterm/pterm" style="text-decoration: none">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-28692-magenta?style=flat-square" alt="Forks"><!-- unittestcount:end -->
</a>

<a href="https://github.com/pterm/pterm/tree/master/_examples/demo" style="text-decoration: none">
<img src="https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg" alt="PTerm">
<p align="center">Show Demo Code</p>
</a>

</p>

<br/>

---

<p align="center">
<strong><a href="https://pterm.sh">PTerm.sh</a></strong>
|
<strong><a href="#-installation">Installation</a></strong>
|
<strong><a href="https://pterm.sh/#/quick-start">Quick Start</a></strong>
|
<strong><a href="https://pterm.sh/#/docs/intro">Documentation</a></strong>
|
<strong><a href="#-examples">Examples</a></strong>
|
<strong><a href="https://github.com/pterm/pterm/discussions?discussions_q=category%3AQ%26A">Q&A</a></strong>
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

## 🥅 Goal of PTerm

PTerm is designed to create a **platform independent way to create beautiful terminal output**. Most modules that want to improve the terminal output do not guarantee platform independence - PTerm does. PTerm follows the **most common methods for displaying color in a terminal**. With PTerm, it is possible to create beautiful output **even in low-level environments**. 

### • 🪀 Easy to use

Our first priority is to keep PTerm as easy to use as possible. With many [examples](#-examples) for each individual component, getting started with PTerm is extremely easy. All components are similar in design and implement interfaces to simplify mixing individual components together.

### • 🤹‍♀️ Cross-Platform

We take special precautions to ensure that PTerm works on as many operating systems and terminals as possible. Whether it's `Windows CMD`, `macOS iTerm2` or in the backend (for example inside a `GitHub Action` or other CI systems), PTerm **guarantees** beautiful output!\
\
*PTerm is actively tested on `Windows`, `Linux (Debian & Ubuntu)` and `macOS`.*

### • ⭐ Main Features

|Feature|Example|Docs|-|Feature|Example|Docs
|---|---|---|---|---|---|---|
|Bar Charts|[Example](#barchart)|[Docs](https://pterm.sh/#/docs/printer/barchart)|-|RGB|[Example](#print-color-rgb)|[Docs](https://pterm.sh/#/docs/printer/rgb)|
|BigText|[Example](#bigtext)|[Docs](https://pterm.sh/#/docs/printer/bigtext)|-|Sections|[Example](#section)|[Docs](https://pterm.sh/#/docs/printer/section)|
|Boxed|[Example](#box)|[Docs](https://pterm.sh/#/docs/printer/box)|-|Spinners|[Example](#spinner)|[Docs](https://pterm.sh/#/docs/printer/spinner)|
|Bullet Lists|[Example](#bulletlist)|[Docs](https://pterm.sh/#/docs/printer/bulletlist)|-|Trees|[Example](#tree)|[Docs](https://pterm.sh/#/docs/printer/tree)|
|Centered|[Example](#center)|[Docs](https://pterm.sh/#/docs/printer/center)|-|Theming|[Example](#theme)|[Docs](https://pterm.sh/#/docs/customizing/theming)|
|Colors|[Example](#print-with-color)|[Docs](https://pterm.sh/#/docs/printer/color)|-|Tables|[Example](#table)|[Docs](https://pterm.sh/#/docs/printer/table)|
|Headers|[Example](#header)|[Docs](https://pterm.sh/#/docs/printer/header)|-|Styles|[Example](#style)|[Docs](https://pterm.sh/#/docs/printer/style)|
|Panels|[Example](#panel)|[Docs](https://pterm.sh/#/docs/printer/panel)|-|Area|[Example](#area)|[Docs](https://pterm.sh/#/docs/printer/area)|
|Paragraphs|[Example](#paragraph)|[Docs](https://pterm.sh/#/docs/printer/paragraph)|-|||
|Prefixes|[Example](#prefix)|[Docs](https://pterm.sh/#/docs/printer/prefix)|-|||
|Progress Bars|[Example](#progressbar)|[Docs](https://pterm.sh/#/docs/printer/progressbar)|-|||

### • 🧪 Well tested

> PTerm has a 100% test coverage, which means that every line of code inside PTerm gets tested automatically

We test PTerm continuously. However, since a human cannot test everything all the time, we have our own test system with which we currently run <!-- unittestcount2:start -->**`28692`**<!-- unittestcount2:end -->
automated tests to ensure that PTerm has no bugs. 

### • ✨ Consistent Colors

PTerm uses the [ANSI color scheme](https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit) which is widely used by terminals to ensure consistent colors in different terminal themes.
If that's not enough, PTerm can be used to access the full RGB color scheme (16 million colors) in terminals that support `TrueColor`.

![ANSI Colors](https://user-images.githubusercontent.com/31022056/96002009-f10c3a80-0e38-11eb-8d90-f3150150599c.png)

### • 📚 Component system

PTerm consists of many components, called `Printers`, which can be used individually or together to generate pretty console output.

### • 🛠 Configurable

PTerm can be used by without any configuration. However, you can easily configure each component with little code, so everyone has the freedom to design their own terminal output.

## ✏ Documentation

To view the official documentation of the latest release, you can go to the automatically generated page of [pkg.go.dev](https://pkg.go.dev/github.com/pterm/pterm#section-documentation) This documentation is very technical and includes every method that can be used in PTerm.

**For an easy start we recommend that you take a look at the [examples section](#-examples).** Here you can see pretty much every feature of PTerm with its source code. The animations of the examples are automatically updated as soon as something changes in PTerm.

Have fun exploring this project 🚀

## 💖 Contributing

If you have found a bug or want to suggest a feature, you can do so [here](https://github.com/pterm/pterm/issues) by opening a new issue.

If you want to contribute to the development of PTerm, you are very welcome to do so. Our contribution guidelines can be found [here](CONTRIBUTING.md).

## 💕 Support

<p align="center">
If you want to support me in further developing my open source projects, you can give me a little tip 😄 
<br/>
Your financial support enables me to focus more on my projects. Thank you very much!
<br/>
<a href="https://www.buymeacoffee.com/marvinjwendt" target="_blank">
<img height="60px" src="https://cdn.buymeacoffee.com/buttons/v2/arial-blue.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" >
</a>
</p>

### 🦸‍♂️ Supporters

|-|User|💸|
|---|---|---|
|![Jens Lauterbach](https://avatars.githubusercontent.com/u/1292368?s=25)|[@jenslauterbach](https://github.com/jenslauterbach)|25$|

## 🧪 Examples

You can find all the examples, with their source code, here: [`./_examples`](./_examples)

<!-- examples:start -->
### area

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/animation.svg)

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

### barchart

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	bars := pterm.Bars{
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

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
	_ = pterm.DefaultBarChart.WithHorizontal().WithBars(bars).Render()
}

```

</details>

### bigtext

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bigtext/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a large text with the LetterStyle from the standard theme.
	// Useful for title screens.
	pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("PTerm")).Render()

	// Print a large text with differently colored letters.
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
}

```

</details>

### box

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Info.Println("This might not be rendered correctly on GitHub,\nbut it will work in a real terminal.\nThis is because GitHub does not use a monospaced font by default for SVGs.")

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

### bulletlist

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print a list with different levels.
	// Useful to generate lists automatically from data.
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}).Render()

	// Convert a text to a list and print it.
	pterm.NewBulletListFromString(`0
 1
  2
   3`, " ").Render()
}

```

</details>

### bulletlist-custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist-custom/animation.svg)

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

### center

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/center/animation.svg)

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

### demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	introScreen()
	clear()
	pseudoApplicationHeader()
	time.Sleep(second)
	installingPseudoList()
	time.Sleep(second * 2)
	pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	installedProgramsSize()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("Tree Printer")
	installedTree()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("TrueColor Support")
	fadeText()
	time.Sleep(second)
	pterm.DefaultSection.Println("Bullet List Printer")
	listPrinter()
}

func installedTree() {
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Go"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
	}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 2, Text: s})
		}
		if s == "pseudo-chrome" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Tabs"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Extensions"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "Refined GitHub"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "GitHub Dark Theme"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Bookmarks"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "PTerm"})
		}
	}

	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(leveledList)).Render()
}

func installingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.UpdateTitle("Installing " + pseudoProgramList[i])
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second / 2)
	}
	p.Stop()
}

func listPrinter() {
	pterm.NewBulletListFromString(`Good bye
 Have a nice day!`, " ").Render()
}

func fadeText() {
	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

	str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)"
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.Info.Println(fadeInfo)
}

func installedProgramsSize() {
	d := pterm.TableData{{"Program Name", "Status", "Size"}}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			d = append(d, []string{s, pterm.LightGreen("pass"), strconv.Itoa(randomInt(7, 200)) + "mb"})
		} else {
			d = append(d, []string{pterm.LightRed(s), pterm.LightRed("fail"), "0mb"})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func pseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func introScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
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

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

```

</details>

### disable-color

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/disable-color/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	pterm.DisableColor()
	introScreen()
	clear()
	pseudoApplicationHeader()
	time.Sleep(second)
	installingPseudoList()
	time.Sleep(second * 2)
	pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	installedProgramsSize()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("Tree Printer")
	installedTree()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("TrueColor Support")
	fadeText()
	time.Sleep(second)
	pterm.DefaultSection.Println("Bullet List Printer")
	listPrinter()
}

func installedTree() {
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Go"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
	}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 2, Text: s})
		}
		if s == "pseudo-chrome" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Tabs"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Extensions"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "Refined GitHub"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "GitHub Dark Theme"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Bookmarks"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "PTerm"})
		}
	}

	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(leveledList)).Render()
}

func installingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.UpdateTitle("Installing " + pseudoProgramList[i])
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second / 2)
	}
	p.Stop()
}

func listPrinter() {
	pterm.NewBulletListFromString(`Good bye
 Have a nice day!`, " ").Render()
}

func fadeText() {
	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

	str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)"
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.Info.Println(fadeInfo)
}

func installedProgramsSize() {
	d := pterm.TableData{{"Program Name", "Status", "Size"}}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			d = append(d, []string{s, pterm.LightGreen("pass"), strconv.Itoa(randomInt(7, 200)) + "mb"})
		} else {
			d = append(d, []string{pterm.LightRed(s), pterm.LightRed("fail"), "0mb"})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func pseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func introScreen() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"PTDP - PTerm Demo Program")

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
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

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

```

</details>

### disable-output

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/disable-output/animation.svg)

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

### disable-styling

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/disable-styling/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	pterm.DisableStyling()
	introScreen()
	clear()
	pseudoApplicationHeader()
	time.Sleep(second)
	installingPseudoList()
	time.Sleep(second * 2)
	pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	installedProgramsSize()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("Tree Printer")
	installedTree()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("TrueColor Support")
	fadeText()
	time.Sleep(second)
	pterm.DefaultSection.Println("Bullet List Printer")
	listPrinter()
}

func installedTree() {
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Go"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
	}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 2, Text: s})
		}
		if s == "pseudo-chrome" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Tabs"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Extensions"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "Refined GitHub"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "GitHub Dark Theme"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Bookmarks"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "PTerm"})
		}
	}

	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(leveledList)).Render()
}

func installingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.UpdateTitle("Installing " + pseudoProgramList[i])
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second / 2)
	}
	p.Stop()
}

func listPrinter() {
	pterm.NewBulletListFromString(`Good bye
 Have a nice day!`, " ").Render()
}

func fadeText() {
	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

	str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)"
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.Info.Println(fadeInfo)
}

func installedProgramsSize() {
	d := pterm.TableData{{"Program Name", "Status", "Size"}}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			d = append(d, []string{s, pterm.LightGreen("pass"), strconv.Itoa(randomInt(7, 200)) + "mb"})
		} else {
			d = append(d, []string{pterm.LightRed(s), pterm.LightRed("fail"), "0mb"})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func pseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func introScreen() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"PTDP - PTerm Demo Program")

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
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

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

```

</details>

### header

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header/animation.svg)

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

### header-custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/header-custom/animation.svg)

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

### override-default-printers

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/override-default-printers/animation.svg)

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

### panel

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/panel/animation.svg)

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

### paragraph

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/paragraph/animation.svg)

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

### paragraph-custom

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/paragraph-custom/animation.svg)

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

### prefix

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/prefix/animation.svg)

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

### print-basic-text

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-basic-text/animation.svg)

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

### print-color-fade

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-color-fade/animation.svg)

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

### print-color-fade-multiple

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-color-fade-multiple/animation.svg)

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

### print-color-rgb

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-color-rgb/animation.svg)

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

### print-with-color

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-with-color/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Print different colored words.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))

	// Print strings with set color.
	pterm.FgBlack.Println("FgBlack")
	pterm.FgRed.Println("FgRed")
	pterm.FgGreen.Println("FgGreen")
	pterm.FgYellow.Println("FgYellow")
	pterm.FgBlue.Println("FgBlue")
	pterm.FgMagenta.Println("FgMagenta")
	pterm.FgCyan.Println("FgCyan")
	pterm.FgWhite.Println("FgWhite")
	pterm.Println() // Print one line space.
	pterm.FgLightRed.Println("FgLightRed")
	pterm.FgLightGreen.Println("FgLightGreen")
	pterm.FgLightYellow.Println("FgLightYellow")
	pterm.FgLightBlue.Println("FgLightBlue")
	pterm.FgLightMagenta.Println("FgLightMagenta")
	pterm.FgLightCyan.Println("FgLightCyan")
	pterm.FgLightWhite.Println("FgLightWhite")
}

```

</details>

### progressbar

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/animation.svg)

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

### section

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/section/animation.svg)

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

### spinner

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/animation.svg)

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

### style

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/style/animation.svg)

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

### table

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/table/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).Render()

	pterm.Println() // Blank line
        
	// Create a table with right alignment.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).WithRightAlignment().Render()
}

```

</details>

### theme

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/theme/animation.svg)

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

### tree

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/tree/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
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
	root := pterm.NewTreeFromLeveledList(leveledList)

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




































































































































