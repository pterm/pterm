package main

import (
	"github.com/pterm/pterm"
)

func main() {

	tis2 := pterm.LvlTreeListItems{
		pterm.LvlTreeListItem{Level: 0, Text: "0.0"},
		pterm.LvlTreeListItem{Level: 1, Text: "0.1"},
		pterm.LvlTreeListItem{Level: 1, Text: "0.2"},
		pterm.LvlTreeListItem{Level: 0, Text: "1.0"},
		pterm.LvlTreeListItem{Level: 0, Text: "2.0"},
		pterm.LvlTreeListItem{Level: 1, Text: "2.1"},
		pterm.LvlTreeListItem{Level: 1, Text: "2.2"},
		pterm.LvlTreeListItem{Level: 0, Text: "3.0"},
		pterm.LvlTreeListItem{Level: 0, Text: "4.0"},
		pterm.LvlTreeListItem{Level: 0, Text: "5.0"},
	}

	tis3 := tis2.ConvertLeveledListToTreeListItems(0, 0)

	pterm.DefaultTreeList.WithItems(tis3).Render()
}
