package main

import "github.com/pterm/pterm"

func main() {
	// Center moves text to the middle of the terminal. By default the text
	// moves as one block, so the lines of a box stay aligned.
	box := pterm.Box.WithTitle("pterm").Sprint("Beautiful terminal output\nwith one line of code")
	pterm.Center.Println(box)
}
