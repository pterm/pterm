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

