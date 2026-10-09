# bulletlist/demo

![Animation](https://vhs.charm.sh/vhs-6ybH1SVurQNm5WaatnNAkF.gif)

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Level controls the indentation depth of each item.
	bulletListItems := []pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}

	pterm.DefaultBulletList.WithItems(bulletListItems).Render()

	// Alternatively, build a list from an indented string. The second
	// argument is the indent unit; one leading space equals one level here.
	text := `0
 1
  2
   3`

	putils.BulletListFromString(text, " ").Render()
}
```
