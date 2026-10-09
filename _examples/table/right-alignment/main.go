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

	// WithRightAlignment right-aligns every cell in the table.
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(tableData).Render()
}
