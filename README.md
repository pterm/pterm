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

## NOTICE

PTerm is currently under development. It is very likely that not all things will remain as they are at the moment. However, PTerm is still functional. The versioning of PTerm follows the SemVer guidelines. Breaking Changes are explicitly mentioned in the changelogs and the version will be increased accordingly. Everybody is welcome to improve PTerm, whether by making suggestions or pull requests. Thanks <3

If you want to wait for a stable release, make sure to star the project and follow it, to get notified when we release v1.0.0 (stable) :rocket:


## Examples

<!-- examples:start -->
### demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.PrintHeader("You can do many things with PTerm")

	time.Sleep(time.Second * 3)

	pterm.Println(pterm.Cyan("Like this header above!"))

	time.Sleep(time.Second * 2)

	clear()

	pterm.PrintSuccess("You can print success messages!")

	time.Sleep(time.Second * 2)

	pterm.PrintInfo("Or infos!")

	time.Sleep(time.Second)

	pterm.PrintError("Or errors!")

	time.Sleep(time.Second)

	pterm.PrintWarning("Or warnings!")

	time.Sleep(time.Second)

	pterm.PrintDescription("Even descriptions can be printed...")

	time.Sleep(time.Second * 2)

	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: []pterm.Color{pterm.FgLightRed, pterm.BgBlue},
		},
	}

	customPrefixPrinter.Println("Or a custom PrefixPrinter can be crafted :)")

	time.Sleep(time.Second * 2)

	pterm.Warning.WithScope("custom-scope").Println("PrefixPrinters also support scopes!")

	time.Sleep(time.Second * 4)

	clear()

	pterm.PrintHeader("Everything can be customized!")

	time.Sleep(time.Second * 2)

	headerStyles := []pterm.Style{
		{pterm.BgGreen},
		{pterm.BgWhite},
		{pterm.BgRed},
		{pterm.BgBlue},
		{pterm.BgYellow},
		{pterm.BgLightMagenta},
	}

	for _, style := range headerStyles {
		clear()
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          5,
		}.Println("Everything can be customized")
		time.Sleep(time.Second / 2)
	}

	for i := 0; i < 10; i++ {
		clear()
		style := headerStyles[len(headerStyles)-1]
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          5 + i,
		}.Println("Everything can be customized")
		time.Sleep(time.Millisecond * 100)
	}

	for i := 0; i < 15; i++ {
		clear()
		style := headerStyles[len(headerStyles)-1]
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          15 - i,
		}.Println("Everything can be customized")
		time.Sleep(time.Millisecond * 100)
	}
}

func clear() {
	print("\033[H\033[2J")
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

### print-header

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-header/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	tick := time.Tick(time.Second * 2)

	// Print with the default HeaderPrinter
	pterm.PrintHeader("This is the default header style")

	<-tick // Wait

	// Print a custom header
	pterm.Header.WithFullWidth().WithTextStyle(pterm.FgDarkGray).WithBackgroundStyle(pterm.BgLightMagenta).Println("Hello, World!")

	<-tick // Wait

	// Create a custom HeaderPrinter
	customHeaderPrinter := pterm.HeaderPrinter{
		TextStyle:       pterm.Style{pterm.FgLightRed},
		BackgroundStyle: pterm.Style{pterm.BgGreen},
		Margin:          15,
	}
	// Use custom Header printer
	customHeaderPrinter.Println("This is a custom header.")
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









































