### table/demo

![Animation](https://vhs.charm.sh/vhs-6GyT4ctanulsK2hMwSXivY.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### table/alternate-row-style

![Animation](https://vhs.charm.sh/vhs-12EtzQW8xBhTNQQ6zHqsUV.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### table/boxed

![Animation](https://vhs.charm.sh/vhs-3Nbskbte7xbT79fb7staX5.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>

### table/from-csv

![Animation](https://vhs.charm.sh/vhs-7xdQkX1Aaj9YV9wPXnXX0l.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	csv := `Firstname,Lastname,Email
Paul,Dean,paul@example.com
Callie,Mckay,callie@example.com
Libby,Camacho,libby@example.com`

	// TableDataFromCSV converts raw CSV into pterm.TableData. The CSV header
	// line ends up as the first row, so WithHasHeader renders it as such.
	pterm.DefaultTable.WithHasHeader().WithData(putils.TableDataFromCSV(csv)).Render()
}
```

</details>

### table/from-structs

![Animation](https://vhs.charm.sh/vhs-7JLw6Gm8B76cG5m18AysEs.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm/putils"
)

// User is a regular struct; no tags or interfaces are needed.
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	users := []User{
		{Name: "Ada Lovelace", Age: 36, Email: "ada@example.com"},
		{Name: "Alan Turing", Age: 41, Email: "alan@example.com"},
		{Name: "Grace Hopper", Age: 85, Email: "grace@example.com"},
	}

	// DefaultTableFromStructSlice fills the default table via reflection: the
	// field names become the first row, so WithHasHeader styles them as the
	// header.
	putils.DefaultTableFromStructSlice(users).WithHasHeader().Render()
}
```

</details>

### table/multiple-lines

![Animation](https://vhs.charm.sh/vhs-roshTIYM84GqS1s50v2rM.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>

### table/right-alignment

![Animation](https://vhs.charm.sh/vhs-1ZkWcvqqmfoUT5tHGxJbX1.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>

