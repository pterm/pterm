package main

import "github.com/pterm/pterm"

func main() {
	// An ordered list numbers its items. Every level counts on its own and
	// starts again under each parent.
	pterm.BulletList.WithOrdered().WithItems([]pterm.BulletListItem{
		{Text: "Install the CLI"},
		{Text: "Log in"},
		{Level: 1, Text: "Open the link that is printed"},
		{Level: 1, Text: "Confirm the code in the browser"},
		{Text: "Deploy your first project"},
	}).Render()
}
