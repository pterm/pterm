package main

import "github.com/pterm/pterm"

func main() {
	// Create a default box with specified padding
	paddedBox := pterm.DefaultBox.WithLeftPadding(4).WithRightPadding(4).WithTopPadding(1).WithBottomPadding(1)

	// Define a title for the box
	title := pterm.LightRed("I'm a box!")

	// Create boxes with the title positioned differently and containing different content
	box1 := paddedBox.WithTitle(title).Sprint("Hello, World!\n      1")                         // Title at default position (top left)
	box2 := paddedBox.WithTitle(title).WithTitleTopCenter().Sprint("Hello, World!\n      2")    // Title at top center
	box3 := paddedBox.WithTitle(title).WithTitleTopRight().Sprint("Hello, World!\n      3")     // Title at top right
	box4 := paddedBox.WithTitle(title).WithTitleBottomRight().Sprint("Hello, World!\n      4")  // Title at bottom right
	box5 := paddedBox.WithTitle(title).WithTitleBottomCenter().Sprint("Hello, World!\n      5") // Title at bottom center
	box6 := paddedBox.WithTitle(title).WithTitleBottomLeft().Sprint("Hello, World!\n      6")   // Title at bottom left
	box7 := paddedBox.WithTitle(title).WithTitleTopLeft().Sprint("Hello, World!\n      7")      // Title at top left

	// Render the boxes in a panel layout
	pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{box1}, {box2}, {box3}},
		{{box4}, {box5}, {box6}},
		{{box7}},
	}).Render()
}
