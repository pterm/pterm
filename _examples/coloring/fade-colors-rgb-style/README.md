# coloring/fade-colors-rgb-style

![Animation](animation.svg)

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
