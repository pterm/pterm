package main

import "github.com/pterm/pterm"

func main() {
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).Render()

	// Create a table with right alignment.
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).WithAlignment(pterm.RightAlignment).Render()
}
