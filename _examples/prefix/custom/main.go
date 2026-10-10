package main

import "github.com/pterm/pterm"

func main() {
	// Status printers are values. With* returns a changed copy, so the
	// package defaults stay untouched.
	deploy := pterm.Success.WithGlyph("↑").WithTone(pterm.ToneAccent)
	deploy.Println("Pushed 3 images to the registry")

	// Store a customized printer to reuse it across the program.
	hint := pterm.Info.WithTone(pterm.ToneMuted)
	hint.Println("Run with --verbose to see every request")
}
