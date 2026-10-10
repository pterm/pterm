package main

import "github.com/pterm/pterm"

func main() {
	// A fixed width lines the rule up with other output.
	pterm.Rule.WithWidth(40).Println("Summary")

	// The title can be muted, and the line can use another glyph and tone.
	pterm.Rule.
		WithWidth(40).
		WithTitleStyle(pterm.ToneMuted.Style()).
		WithGlyph("=").
		WithTone(pterm.ToneMuted).
		Println("Details")
}
