package main

import (
	"atomicgo.dev/color"

	"github.com/pterm/pterm"
)

func main() {
	// A Style combines a foreground tone, an optional background tone and
	// text attributes. Tones follow the active theme.
	title := pterm.NewStyle(pterm.TonePrimary, color.Bold)
	label := pterm.NewStyle(pterm.ToneText).On(pterm.ToneBorder)
	link := pterm.NewStyle(pterm.ToneInfo, color.Underline)

	pterm.Println(title.Sprint("Release notes"))
	pterm.Println(label.Sprint(" v0.13.0 "), pterm.ToneMuted.Sprint("2026-10-10"))
	pterm.Println("Read more at", link.Sprint("https://pterm.sh"))

	// With adds attributes to a copy, so the base style stays untouched.
	pterm.Println(title.With(color.Italic).Sprint("Thanks to all contributors"))
}
