package main

import "github.com/pterm/pterm"

func main() {
	// A gradient can run through any number of tones.
	pterm.BigText.WithGradient(pterm.ToneInfo, pterm.TonePrimary, pterm.ToneError).Println("v1.0")

	pterm.Println()

	// The offset shifts the gradient by a share of the text width. It
	// bounces at the ends, so stepping it from 0 to 2 in a loop sweeps the
	// colors across the text and back without a jump.
	pterm.BigText.WithGradientOffset(0.5).Println("PTerm")
}
