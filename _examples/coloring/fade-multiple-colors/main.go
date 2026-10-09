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
