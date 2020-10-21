package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "Level 0"},
		{Level: 1, Text: "Level 1"},
		{Level: 2, Text: "Level 2"},
	}).Render()

	pterm.NewListFromString(`0
 1
  2
   3`, " ").Render()
}
