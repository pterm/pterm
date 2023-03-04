package main

import "github.com/pterm/pterm"

func main() {
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com", "just a test, hey!"},
	}).Render()

	pterm.Println() // Blank line

	// Create a table with multiple lines in a row.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}).Render()
}
