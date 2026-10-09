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
