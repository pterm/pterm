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
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}

```

</details>

