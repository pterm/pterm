package main

import "github.com/pterm/pterm"

func main() {
	// Any tone of the theme can be used. Keep tones for meaning: muted
	// text for hints, accent for values worth noticing.
	pterm.Text.WithTone(pterm.ToneMuted).Println("Hints and metadata are muted.")
	pterm.Text.WithTone(pterm.ToneAccent).Println("Values stand out in the accent color.")

	// WithNoWrap keeps lines exactly as they are, for preformatted output.
	pterm.Text.WithNoWrap().Println("id    name     status\n42    worker   running")
}
