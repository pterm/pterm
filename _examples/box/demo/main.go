package main

import "github.com/pterm/pterm"

func main() {
	// A box frames text that belongs together. The title sits in the
	// border, so the frame costs no extra line.
	pterm.Box.WithTitle("Result").Println("3 files changed\n128 insertions, 41 deletions")

	// Boxes are text printers, so their output can go into other printers,
	// even into another box.
	inner := pterm.Box.WithTitle("Cache").Sprint("12 entries\n4.1 MB")
	pterm.Box.WithTitle("Status").Println(inner)
}
