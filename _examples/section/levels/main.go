package main

import "github.com/pterm/pterm"

func main() {
	// Deeper levels get quieter: level 2 is bold text, level 3 is muted.
	pterm.Section.Println("Packages")
	pterm.Section.WithLevel(2).Println("Runtime")
	pterm.Text.Println("atomicgo.dev/color, atomicgo.dev/cursor")
	pterm.Section.WithLevel(3).Println("Indirect")
	pterm.Text.Println("golang.org/x/term")

	// A prefix makes the depth visible, like headings in Markdown.
	pterm.Section.WithPrefix("#").WithLevel(2).WithBottomPadding(1).Println("Development")
	pterm.Text.Println("github.com/pterm/pterm/internal/vt")
}
