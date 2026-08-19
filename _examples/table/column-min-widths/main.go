package main

import "github.com/pterm/pterm"

func main() {
	// Give every table the same minimum column widths so they line up when
	// printed one after another, even though their content differs in length.
	widths := []int{14, 8}

	pterm.DefaultTable.WithHasHeader().WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Alice", "1"},
		{"Bob", "22"},
	}).Render()

	pterm.Println()

	pterm.DefaultTable.WithHasHeader().WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Christopher", "3"},
		{"Dan", "444"},
	}).Render()
}
