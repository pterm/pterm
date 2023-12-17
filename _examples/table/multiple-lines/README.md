# table/multiple-lines

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	data := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Create and render the table.
	// The options are chained in a single line for simplicity.
	// The table has a header, a row separator, and a header row separator.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(data).Render()
}

```
