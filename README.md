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

</p>

<a href="https://github.com/pterm/pterm/">
<img src="https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg" alt="Dops">
</a>

---

## Examples

<!-- examples:start -->
### demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/demo/animation.svg)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {

	for _, p := range pterm.AllPrinters {
		p.Println("This is the default ", p.Prefix.Text, " printer.")
		p.WithScope("scope").Println("This is the ", p.Prefix.Text, " printer with a scope.")
		p.WithScope("custom", pterm.New(pterm.FgLightMagenta, pterm.Bold, pterm.BgWhite)).Println("This is the ", p.Prefix.Text, " printer with a custom scope style.")
		time.Sleep(time.Second)
		pterm.Println()
	}

	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: []pterm.Color{pterm.FgLightRed, pterm.BgBlue},
		},
	}

	customPrefixPrinter.Println("This is a custom PrefixPrinter :)")

}

```

### override-default-printer

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/override-default-printer/animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {

	pterm.ErrorPrinter.Println("This is the default ErrorPrinter")

	pterm.ErrorPrinter.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.Style{pterm.BgCyan, pterm.FgRed},
	}

	pterm.ErrorPrinter.Println("This is the default ErrorPrinter after the prefix was overridden")
}

```

### print-with-color

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-with-color/animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Simple Println with different colored words.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
}

```

<!-- examples:end -->





