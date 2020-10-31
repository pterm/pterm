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
		pterm.LvlTreeListItem{Level: 2, Text: "2.2.1"},
		pterm.LvlTreeListItem{Level: 1, Text: "2.3"},
		pterm.LvlTreeListItem{Level: 2, Text: "2.3.1"},
		pterm.LvlTreeListItem{Level: 3, Text: "2.3.1.1"},
		pterm.LvlTreeListItem{Level: 4, Text: "2.3.1.1.1"},
		pterm.LvlTreeListItem{Level: 0, Text: "3.0"},
		pterm.LvlTreeListItem{Level: 0, Text: "4.0"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.0"},
		pterm.LvlTreeListItem{Level: 2, Text: "4.0.1"},
		pterm.LvlTreeListItem{Level: 3, Text: "4.0.1.1"},
		pterm.LvlTreeListItem{Level: 4, Text: "4.0.1.1.1"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.1"},
		pterm.LvlTreeListItem{Level: 2, Text: "4.1.1"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.2"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.3"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.4"},
		pterm.LvlTreeListItem{Level: 1, Text: "4.5"},
		pterm.LvlTreeListItem{Level: 0, Text: "5.0"},
	}

	tis3 := pterm.ConvertLeveledListToTreeListItems(tis2)

	pterm.DefaultTreeList.WithItems(tis3).Render()
}
