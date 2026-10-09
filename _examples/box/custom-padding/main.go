package main

import "github.com/pterm/pterm"

func main() {
	// Padding adds empty space between the box border and its content,
	// configurable per side.
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Hello, World!")
}
