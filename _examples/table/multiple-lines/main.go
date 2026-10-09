package main

import "github.com/pterm/pterm"

func main() {
	// Cells may contain newlines; a row grows to fit its tallest cell.
	data := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Row separators keep multi-line rows visually apart.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(data).Render()
}
