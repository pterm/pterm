package main

import "github.com/pterm/pterm"

func main() {
	// The first row becomes the header via WithHasHeader.
	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// WithBoxed draws a box around the whole table.
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
}
