<!--suppress HtmlDeprecatedAttribute -->

<h1 align="center">:computer: PTerm | Pretty Terminal</h1>
<p align="center">A golang module to print pretty text</p>

<p align="center">

<a href="https://github.com/pterm/pterm/releases">
<img src="https://img.shields.io/github/v/release/pterm/pterm?style=flat-square" alt="Latest Release">
</a>

<a href="https://github.com/pterm/pterm/stargazers">
<img src="https://img.shields.io/github/stars/pterm/pterm.svg?style=flat-square" alt="Stars">
</a>

<a href="https://github.com/pterm/pterm/fork">
<img src="https://img.shields.io/github/forks/pterm/pterm.svg?style=flat-square" alt="Forks">
</a>

<a href="https://github.com/pterm/pterm/issues">
<img src="https://img.shields.io/github/issues/pterm/pterm.svg?style=flat-square" alt="Issues">
</a>

<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a href="https://github.com/dops-cli/dops/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

<a href="https://pkg.go.dev/github.com/pterm/pterm">
<img src="https://pkg.go.dev/badge/github.com/pterm/pterm" alt="PTerm Documentation"/>
</a>

<a href="https://github.com/pterm/pterm/">
<img src="https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg" alt="Dops">
</a>

</p>

---

## :warning: NOTICE

PTerm is currently under development. It is very likely that not all things will remain as they are at the moment. However, PTerm is still functional. The versioning of PTerm follows the SemVer guidelines. Breaking Changes are explicitly mentioned in the changelogs and the version will be increased accordingly. Everybody is welcome to improve PTerm, whether by making suggestions or pull requests. Thanks <3

If you want to wait for a stable release, make sure to star the project and follow it, to get notified when we release v1.0.0 (stable) :rocket:

## :package: Installation

To make PTerm available in your project, you can run the following command.\
Make sure to run this command inside your project, when you're using go modules :wink:

```sh
go get github.com/pterm/pterm
```

## :pencil: Documentation

To view the official documentation of the latest release, you can go to the automatically generated page of [pkg.go.dev](https://pkg.go.dev/github.com/pterm/pterm) This documentation is very technical and includes every method that can be used in PTerm.

**For an easy start we recommend that you take a look at the [examples section](#test_tube-examples).** Here you can see pretty much every feature of PTerm with its source code. The animations of the examples are automatically updated as soon as something changes in PTerm.

Have fun exploring this project :rocket:

## :heart: Contributing

If you have found a bug or want to suggest a feature, you can do so [here](https://github.com/pterm/pterm/issues) by opening a new issue.

If you want to contribute to the development of PTerm, you are very welcome to do so. Our contribution guidelines can be found [here](CONTRIBUTING.md).

## :test_tube: Examples

<!-- examples:start -->
### demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var (
	pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
		"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")
)

func main() {
	// Change this to time.Millisecond*200 to speed up the demo.
	// Useful when debugging.
	const second = time.Second

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

	introSpinner := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
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

	clear()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")

	time.Sleep(second)

	setupSpinner := pterm.DefaultSpinner.Start("Fetching pseudo install list...")
	time.Sleep(second * 4)
	setupSpinner.Success()

	p := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Downloading stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Downloading " + pseudoProgramList[i]
		pterm.Success.Println("Downloading " + pseudoProgramList[i])
		p.Increment()
		time.Sleep(time.Millisecond * 500)
	}
	pterm.Success.Println("Downloaded all pseudo programs!")

	pterm.Println()
	pterm.Info.Println("Installing pseudo programs")

	p = pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Installing " + pseudoProgramList[i]
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second)
	}
}

func clear() {
	print("\033[H\033[2J")
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
	pterm.DefaultHeader.Println("This is the default header!")
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
				WithMargin(15).                                    // Set Margin to 15
				WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)). // Set BackgroundStyle to Cyan
				WithTextStyle(pterm.NewStyle(pterm.FgBlack)).      // Set TextStyle to Black
				Println("This is a custom header!")                // Print header
	// Instead of printing the header you can set it to a variable.
	// You can then reuse your custom header.

	// Making a completely new HeaderPrinter
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	newHeader.Println("This is a custom header!")

}

```

</details>

### override-default-printer

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/override-default-printer/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Error.Println("This is the default Error")

	pterm.Error.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.Style{pterm.BgCyan, pterm.FgRed},
	}

	pterm.Error.Println("This is the default Error after the prefix was overridden")
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
	// Simple Println with different colored words.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))
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

var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

var vki int

func main() {
	p := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		p.Title = "Downloading " + fakeInstallList[vki]
		pterm.Success.Println("Downloading " + fakeInstallList[vki])
		vki++
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}

	pterm.Success.Println("Finished downloading!")

	time.Sleep(time.Second * 5)
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
	pterm.DefaultSection.Println("This is a section!")
	pterm.Info.Println("And here is some text.\nThis text could be anything.\nBasically it's just a placeholder")
	pterm.DefaultSection.Println("This is another section!")
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
	spinnerSuccess := pterm.DefaultSpinner.Start("Doing something important... (will succeed)")

	time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	spinnerSuccess.Success()

	spinnerWarning := pterm.DefaultSpinner.Start("Doing something important... (will warn)")

	time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	spinnerWarning.Warning()

	spinnerFail := pterm.DefaultSpinner.Start("Doing something important... (will fail)")

	time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	spinnerFail.Fail()

	spinnerLiveText := pterm.DefaultSpinner.Start("Doing a lot of stuff...")

	time.Sleep(time.Second * 2)

	spinnerLiveText.UpdateText("It's really much")

	time.Sleep(time.Second * 2)

	spinnerLiveText.UpdateText("We're nearly done!")

	time.Sleep(time.Second * 2)

	spinnerLiveText.Success("Finally!")
}

```

</details>

<!-- examples:end -->












































































