# coloring/fade-colors-rgb-style

![Animation](animation.svg)

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
