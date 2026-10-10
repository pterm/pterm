package main

import "github.com/pterm/pterm"

func main() {
	// The same minimum widths keep two tables aligned when they are printed
	// one after another, even though their content differs in length.
	widths := []int{14, 8}

	pterm.Table.WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Alice", "1"},
		{"Bob", "22"},
	}).Render()

	pterm.Println()

	pterm.Table.WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Christopher", "3"},
		{"Dan", "444"},
	}).Render()
}
