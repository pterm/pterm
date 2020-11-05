package main

import "github.com/pterm/pterm"

func main() {
	// Declare panels in a two dimensional grid system.
	panels := pterm.Panels{
		{{Data: "This is the first panel"}, {Data: pterm.DefaultHeader.Sprint("Hello, World!")}, {Data: "This\npanel\ncontains\nmultiple\nlines"}},
		{{Data: pterm.Red("This is another\npanel line")}, {Data: "This is the second panel\nwith a new line"}},
	}

	// Print panels.
	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}
