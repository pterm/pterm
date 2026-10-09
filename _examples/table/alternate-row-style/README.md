# table/alternate-row-style

![Animation](https://vhs.charm.sh/vhs-12EtzQW8xBhTNQQ6zHqsUV.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// WithAlternateRowStyle applies this style to every second data row,
	// which makes wide tables easier to scan.
	alternateStyle := pterm.NewStyle(pterm.BgDarkGray)

	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
}
```
