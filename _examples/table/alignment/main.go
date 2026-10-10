package main

import "github.com/pterm/pterm"

func main() {
	data := pterm.TableData{
		{"Package", "Version", "Downloads"},
		{"pterm", "v0.13.0", "1,204,331"},
		{"color", "v0.2.0", "88,412"},
		{"keyboard", "v0.3.1", "41,007"},
	}

	// Numbers are right-aligned on their own. Columns can still be aligned
	// by hand, here the version column is centered.
	pterm.Table.WithData(data).WithColumnAlignment(1, pterm.AlignCenter).Render()

	pterm.Println()

	// WithRightAlignment aligns every column to the right.
	pterm.Table.WithData(data).WithRightAlignment().Render()
}
