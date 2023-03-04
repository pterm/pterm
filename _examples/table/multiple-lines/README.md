# table/multiple-lines

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a table with multiple lines in a row and set a row separator.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}).Render()
}

```
