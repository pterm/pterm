package main

import "github.com/pterm/pterm"

func main() {
	// Letters with a style of their own keep it; the rest follow the
	// printer's gradient.
	pterm.BigText.WithLetters(
		pterm.LettersFromString("P", pterm.ToneAccent.Style()),
		pterm.LettersFromString("Term"),
	).Render()
}
