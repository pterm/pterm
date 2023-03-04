# table/right-alignment

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a fork of the default table, fill it with data and print it.
	// Data can also be generated and inserted later.
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com", "just a test, hey!"},
	}).Render()
}

```
