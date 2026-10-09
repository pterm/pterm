# table/demo

![Animation](https://vhs.charm.sh/vhs-6GyT4ctanulsK2hMwSXivY.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// With WithHasHeader the first row is styled as the header. The CJK
	// characters are measured by display width, so the columns stay aligned.
	tableData1 := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

	pterm.Println()

	// Cells may contain newlines; a row grows to fit its tallest cell.
	tableData2 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	pterm.DefaultTable.WithHasHeader().WithData(tableData2).Render()
}
```
