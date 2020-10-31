package main

import (
	"github.com/pterm/pterm"
)

func main() {

	tis2 := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "0.0"},
		pterm.LeveledListItem{Level: 1, Text: "0.1"},
		pterm.LeveledListItem{Level: 1, Text: "0.2"},
		pterm.LeveledListItem{Level: 0, Text: "1.0"},
		pterm.LeveledListItem{Level: 0, Text: "2.0"},
		pterm.LeveledListItem{Level: 1, Text: "2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.2"},
		pterm.LeveledListItem{Level: 2, Text: "2.2.1"},
		pterm.LeveledListItem{Level: 1, Text: "2.3"},
		pterm.LeveledListItem{Level: 2, Text: "2.3.1"},
		pterm.LeveledListItem{Level: 3, Text: "2.3.1.1"},
		pterm.LeveledListItem{Level: 4, Text: "2.3.1.1.1"},
		pterm.LeveledListItem{Level: 0, Text: "3.0"},
		pterm.LeveledListItem{Level: 0, Text: "4.0"},
		pterm.LeveledListItem{Level: 1, Text: "4.0"},
		pterm.LeveledListItem{Level: 2, Text: "4.0.1"},
		pterm.LeveledListItem{Level: 3, Text: "4.0.1.1"},
		pterm.LeveledListItem{Level: 4, Text: "4.0.1.1.1"},
		pterm.LeveledListItem{Level: 1, Text: "4.1"},
		pterm.LeveledListItem{Level: 2, Text: "4.1.1"},
		pterm.LeveledListItem{Level: 1, Text: "4.2"},
		pterm.LeveledListItem{Level: 1, Text: "4.3"},
		pterm.LeveledListItem{Level: 0, Text: "5.0"},
	}

	tis3 := pterm.NewTreeFromLeveledList(tis2)

	pterm.DefaultTree.WithRoot(tis3).Render()
}
