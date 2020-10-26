package main

import (
	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("Please note that RGB colors only work in Terminals which support TrueColor.")

	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)
	to2 := pterm.NewRGB(255, 0, 0)
	to3 := pterm.NewRGB(0, 255, 0)
	to4 := pterm.NewRGB(255, 255, 255)
	for i := 0; i < pterm.GetTerminalHeight()-1; i++ {
		from.Fade(0, float32(pterm.GetTerminalHeight()), float32(i), to, to2, to3, to4).Println("Hello, World!")
	}
}
