package main

import (
	"github.com/pterm/pterm"
)

func main() {
	rgb1 := pterm.RGB{R: 187, G: 80, B: 0}
	rgb2 := pterm.RGB{R: 0, G: 50, B: 123}
	pterm.NewRGBStyle(rgb1, rgb2).Println("This text is not styled.")
	pterm.NewRGBStyle(rgb1, rgb2).AddOptions(pterm.Bold).Println("This text is bold.")
	pterm.NewRGBStyle(rgb1, rgb2).AddOptions(pterm.Italic).Println("This text is italic.")
}
