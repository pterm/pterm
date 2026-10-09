### coloring/demo

![Animation](https://vhs.charm.sh/vhs-56kMr16EfHRqxQbkxLF5sM.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Show every built-in foreground (Fg*) and background (Bg*) color in one
	// table. The "Light" variants map to the bright ANSI colors.
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

	// Shorthand functions like pterm.Red return colored strings that can be
	// concatenated, even nested inside each other.
	pterm.Println(pterm.Red("Hello, ") + pterm.Green("World") + pterm.Cyan("!"))
	pterm.Println(pterm.Red("Even " + pterm.Cyan("nested ") + pterm.Green("colors ") + "are supported!"))

	pterm.Println()

	// NewStyle combines multiple attributes into a reusable style.
	style := pterm.NewStyle(pterm.BgRed, pterm.FgLightGreen, pterm.Bold)
	style.Println("This text uses a style and is bold and light green with a red background!")
}
```

</details>

### coloring/disable-output

![Animation](https://vhs.charm.sh/vhs-6JuPDJLeFeUCl3biL6qiP0.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// DisableOutput silences all PTerm printers globally until EnableOutput
	// is called. Iterations 5-9 below produce no output at all.
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

![Animation](https://vhs.charm.sh/vhs-6uuzyL2JxFTfYSemtD7OiA.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// RGB colors need a TrueColor terminal; on anything less PTerm
	// downsamples them to the closest supported color.
	pterm.Info.Println("RGB colors only work in Terminals which support TrueColor.")

	startColor := pterm.NewRGB(0, 255, 255) // cyan
	endColor := pterm.NewRGB(255, 0, 255)   // magenta

	// Spread the gradient over the visible terminal height, one line per step.
	terminalHeight := pterm.GetTerminalHeight()

	for i := 0; i < terminalHeight-2; i++ {
		// Fade interpolates between the two colors; the factor 0..1 selects
		// the position on the gradient.
		fadeFactor := float32(i) / float32(terminalHeight-2)
		currentColor := startColor.Fade(0, 1, fadeFactor, endColor)

		currentColor.Println("Hello, World!")
	}
}
```

</details>

### coloring/fade-colors-rgb-style

![Animation](https://vhs.charm.sh/vhs-2YMnxtKIWCFDzve44qyKuw.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"

	"github.com/pterm/pterm"
)

// Demonstrates RGBStyle: fading the foreground and background independently,
// plus adding options like Bold or Italic on top of an RGB gradient.
// RGB colors need a TrueColor terminal to show up as smooth gradients.
func main() {
	white := pterm.NewRGB(255, 255, 255)
	grey := pterm.NewRGB(128, 128, 128)
	black := pterm.NewRGB(0, 0, 0)
	red := pterm.NewRGB(255, 0, 0)
	purple := pterm.NewRGB(255, 0, 255)
	green := pterm.NewRGB(0, 255, 0)

	str1 := "RGB colors only work in Terminals which support TrueColor."
	str2 := "The background and foreground colors can be customized individually."
	str3 := "Styles can also be applied. For example: Bold or Italic."

	printFadedString(str1, white, purple, grey, black)
	printFadedString(str2, black, purple, red, red)
	printStyledString(str3, white, green, red, black)
}

// printFadedString fades the foreground from fgStart to fgEnd and the
// background from bgStart to bgEnd across the string, one character at a time.
func printFadedString(str string, fgStart, fgEnd, bgStart, bgEnd pterm.RGB) {
	strs := strings.Split(str, "")
	var result string
	for i := 0; i < len(str); i++ {
		style := pterm.NewRGBStyle(fgStart.Fade(0, float32(len(str)), float32(i), fgEnd), bgStart.Fade(0, float32(len(str)), float32(i), bgEnd))
		result += style.Sprint(strs[i])
	}
	pterm.Println(result)
}

// printStyledString does the same fade, but additionally renders the words
// "Bold" and "Italic" in their respective style when they appear in the text.
func printStyledString(str string, fgStart, fgEnd, bgStart, bgEnd pterm.RGB) {
	strs := strings.Split(str, "")
	var result string
	boldStr := strings.Split("Bold", "")
	italicStr := strings.Split("Italic", "")
	bold, italic := 0, 0
	for i := 0; i < len(str); i++ {
		style := pterm.NewRGBStyle(fgStart.Fade(0, float32(len(str)), float32(i), fgEnd), bgStart.Fade(0, float32(len(str)), float32(i), bgEnd))
		// While inside the word "Bold" or "Italic", add the matching option.
		if bold < len(boldStr) && i+len(boldStr)-bold <= len(strs) && strings.Join(strs[i:i+len(boldStr)-bold], "") == strings.Join(boldStr[bold:], "") {
			style = style.AddOptions(pterm.Bold)
			bold++
		} else if italic < len(italicStr) && i+len(italicStr)-italic < len(strs) && strings.Join(strs[i:i+len(italicStr)-italic], "") == strings.Join(italicStr[italic:], "") {
			style = style.AddOptions(pterm.Italic)
			italic++
		}
		result += style.Sprint(strs[i])
	}
	pterm.Println(result)
}
```

</details>

### coloring/fade-multiple-colors

![Animation](https://vhs.charm.sh/vhs-3bXKSqrMcEBlnZIrBhGbHE.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	// Fade accepts multiple target colors, so a gradient can pass through
	// several points instead of just blending between two.
	startColor := pterm.NewRGB(0, 255, 255)
	firstPoint := pterm.NewRGB(255, 0, 255)
	secondPoint := pterm.NewRGB(255, 0, 0)
	thirdPoint := pterm.NewRGB(0, 255, 0)
	endColor := pterm.NewRGB(255, 255, 255)

	str := "RGB colors only work in Terminals which support TrueColor."
	strs := strings.Split(str, "")

	// Fade a single line horizontally, character by character.
	var fadeInfo string
	for i := 0; i < len(str); i++ {
		fadeInfo += startColor.Fade(0, float32(len(str)), float32(i), firstPoint).Sprint(strs[i])
	}

	pterm.Info.Println(fadeInfo)

	terminalHeight := pterm.GetTerminalHeight()

	// Fade vertically over the visible terminal height, passing through all
	// four gradient points from top to bottom.
	for i := 0; i < terminalHeight-2; i++ {
		startColor.Fade(0, float32(terminalHeight-2), float32(i), firstPoint, secondPoint, thirdPoint, endColor).Println("Hello, World!")
	}
}
```

</details>

### coloring/override-default-printers

![Animation](https://vhs.charm.sh/vhs-5uqfqe2Y4nIt6DAahZxgrR.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Error.Println("This is the default Error")

	// The default printers are package-level variables, so their fields can
	// be changed directly. This swaps the Error prefix text and style for
	// every subsequent pterm.Error call in the program.
	pterm.Error.Prefix = pterm.Prefix{Text: "OVERRIDE", Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed)}

	pterm.Error.Println("This is the default Error after the prefix was overridden")
}
```

</details>

### coloring/print-color-rgb

![Animation](https://vhs.charm.sh/vhs-4Iok4H7S6qYWGRtf5H2dTf.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// NewRGB creates a TrueColor foreground color; every RGB value can be
	// used as a printer directly.
	pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")
	pterm.NewRGB(15, 199, 209).Println("This text is printed with a custom RGB!")

	// Passing true as the optional last argument makes the color apply to
	// the background instead of the text.
	pterm.NewRGB(201, 144, 30, true).Println("This text is printed with a custom RGB background!")
}
```

</details>

### coloring/print-color-rgb-style

![Animation](https://vhs.charm.sh/vhs-78d44TVBSv37OO95sBQ22Z.gif)

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

	// NewRGBStyle pairs a TrueColor foreground with a background color.
	rgbStyle := pterm.NewRGBStyle(foregroundRGB, backgroundRGB)

	rgbStyle.Println("This text is not styled.")

	// AddOptions returns a new style, so the bold and italic lines below are
	// independent of each other and of rgbStyle.
	rgbStyle.AddOptions(pterm.Bold).Println("This text is bold.")
	rgbStyle.AddOptions(pterm.Italic).Println("This text is italic.")
}
```

</details>

