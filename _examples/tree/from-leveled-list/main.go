package main

import "github.com/pterm/pterm"

func main() {
	// A LeveledList is a flat alternative to nesting TreeNodes by hand: each
	// item states its own depth.
	list := pterm.LeveledList{
		{Level: 0, Text: "C:"},
		{Level: 1, Text: "Users"},
		{Level: 1, Text: "Windows"},
		{Level: 1, Text: "Programs"},
		{Level: 0, Text: "D:"},
		{Level: 1, Text: "Movies"},
		{Level: 1, Text: "Games"},
		{Level: 2, Text: "Shooter"},
		{Level: 3, Text: "Battlefield"},
		{Level: 0, Text: "E:"},
		{Level: 1, Text: "dev"},
		{Level: 2, Text: "pterm"},
	}

	root := list.Tree()
	root.Text = "Computer"

	pterm.Tree.WithRoot(root).Render()
}
