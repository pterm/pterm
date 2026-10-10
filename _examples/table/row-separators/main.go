package main

import "github.com/pterm/pterm"

func main() {
	// Row separators work in boxed tables too, where they join the frame.
	pterm.Table.WithBoxed().WithRowSeparators().WithData(pterm.TableData{
		{"Region", "Nodes", "Load"},
		{"eu-central", "12", "64%"},
		{"us-east", "30", "81%"},
		{"ap-south", "6", "23%"},
	}).Render()
}
