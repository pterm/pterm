package main

import (
	"github.com/pterm/pterm"
)

func main() {

	var is []pterm.ListItem

	is = append(is, pterm.DefaultListItem.WithLevel(0).WithText("0"))
	is = append(is, pterm.DefaultListItem.WithLevel(1).WithText("1"))
	is = append(is, pterm.DefaultListItem.WithLevel(2).WithText("2"))
	is = append(is, pterm.DefaultListItem.WithLevel(3).WithText("3"))
	is = append(is, pterm.DefaultListItem.WithLevel(4).WithText("4"))
	is = append(is, pterm.DefaultListItem.WithLevel(5).WithText("5"))

	pterm.DefaultBulletList.WithItems(is).Render()

	pterm.NewListFromString(`0
 1
  2
   3
    4
     5`, " ").Render()

	pterm.NewListFromStrings([]string{"0", " 1", "  2", "   3", "    4", "     5"}, " ").Render()

	pterm.DefaultBulletList.WithItems([]pterm.ListItem{
		{Level: 0, Text: "Hello World"},
		{Level: 1, Text: "Hello World"},
		{Level: 1, Text: "Hello World"},
		{Level: 2, Text: "Hello World"},
		{Level: 3, Text: "Hello World"},
		{Level: 4, Text: "Hello World"},
		{Level: 5, Text: "Hello World"},
		{Level: 6, Text: "Hello World"},
	}).Render()
}
