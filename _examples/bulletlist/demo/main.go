package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Define a list of bullet list items with different levels.
	bulletListItems := []pterm.BulletListItem{
		{Level: 0, Text: "Level 0"}, // Level 0 item
		{Level: 1, Text: "Level 1"}, // Level 1 item
		{Level: 2, Text: "Level 2"}, // Level 2 item
	}

	// Use the default bullet list style to render the list items.
	pterm.DefaultBulletList.WithItems(bulletListItems).Render()

	// Define a string with different levels of indentation.
	text := `0
 1
  2
   3`

	// Convert the indented string to a bullet list and render it.
	putils.BulletListFromString(text, " ").Render()
}
