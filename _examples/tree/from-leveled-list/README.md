# tree/from-leveled-list

![Animation](https://vhs.charm.sh/vhs-5m4lW2DbvrN4pgynZUdVvO.gif)

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// A LeveledList is a flat alternative to nesting TreeNodes by hand: each
	// entry states its own depth.
	leveledList := pterm.LeveledList{
		{Level: 0, Text: "C:"},
		{Level: 1, Text: "Users"},
		{Level: 1, Text: "Windows"},
		{Level: 1, Text: "Programs"},
		{Level: 1, Text: "Programs(x86)"},
		{Level: 1, Text: "dev"},
		{Level: 0, Text: "D:"},
		{Level: 0, Text: "E:"},
		{Level: 1, Text: "Movies"},
		{Level: 1, Text: "Music"},
		{Level: 2, Text: "LinkinPark"},
		{Level: 1, Text: "Games"},
		{Level: 2, Text: "Shooter"},
		{Level: 3, Text: "CallOfDuty"},
		{Level: 3, Text: "CS:GO"},
		{Level: 3, Text: "Battlefield"},
		{Level: 4, Text: "Battlefield 1"},
		{Level: 4, Text: "Battlefield 2"},
		{Level: 0, Text: "F:"},
		{Level: 1, Text: "dev"},
		{Level: 2, Text: "dops"},
		{Level: 2, Text: "PTerm"},
	}

	// TreeFromLeveledList converts the list into a TreeNode; the returned
	// root just needs a name.
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = "Computer"

	pterm.DefaultTree.WithRoot(root).Render()
}
```
