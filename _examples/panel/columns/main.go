package main

import "github.com/pterm/pterm"

func main() {
	// Without frames, panels are plain columns. Their content can come from
	// other printers, such as these bullet lists.
	tools, _ := pterm.BulletList.WithTexts("npm", "bun", "pip").Srender()
	stores, _ := pterm.BulletList.WithTexts("Docker images", "Go modules").Srender()

	pterm.Panel.WithPadding(6).WithPanels(pterm.Panels{{
		{Title: "Package managers", Data: tools},
		{Title: "Stores", Data: stores},
	}}).Render()
}
