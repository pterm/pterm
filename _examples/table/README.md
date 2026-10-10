### table/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The first row is the header. Sizes and numbers are detected and
	// right-aligned, so the SIZE and FILES columns line up on their units.
	pterm.Table.WithData(pterm.TableData{
		{"Cache", "State", "Size", "Files", "Path"},
		{"npm", "stale", "1.20 GB", "48,211", "~/.npm/_cacache"},
		{"bun", "fresh", "393 MB", "9,874", "~/.bun/install/cache"},
		{"go", "fresh", "2.84 GB", "112,530", "~/go/pkg/mod"},
		{"pip", "stale", "87 MB", "1,302", "~/.cache/pip"},
		{"cargo", "fresh", "640 MB", "20,447", "~/.cargo/registry"},
	}).Render()
}
```

</details>

### table/alignment

![Animation](alignment/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	data := pterm.TableData{
		{"Package", "Version", "Downloads"},
		{"pterm", "v0.13.0", "1,204,331"},
		{"color", "v0.2.0", "88,412"},
		{"keyboard", "v0.3.1", "41,007"},
	}

	// Numbers are right-aligned on their own. Columns can still be aligned
	// by hand, here the version column is centered.
	pterm.Table.WithData(data).WithColumnAlignment(1, pterm.AlignCenter).Render()

	pterm.Println()

	// WithRightAlignment aligns every column to the right.
	pterm.Table.WithData(data).WithRightAlignment().Render()
}
```

</details>

### table/boxed

![Animation](boxed/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A boxed table gets a rounded frame with its title inset into the top
	// edge. Use it when the table has to stand apart from the output around it.
	pterm.Table.WithBoxed().WithTitle("Services").WithData(pterm.TableData{
		{"Service", "Region", "Status", "Latency"},
		{"api", "eu-west-1", "running", "42 ms"},
		{"worker", "eu-west-1", "running", "118 ms"},
		{"billing", "us-east-1", "degraded", "1.4 s"},
	}).Render()
}
```

</details>

### table/column-min-widths

![Animation](column-min-widths/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The same minimum widths keep two tables aligned when they are printed
	// one after another, even though their content differs in length.
	widths := []int{14, 8}

	pterm.Table.WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Alice", "1"},
		{"Bob", "22"},
	}).Render()

	pterm.Println()

	pterm.Table.WithColumnMinWidths(widths...).WithData(pterm.TableData{
		{"Name", "Score"},
		{"Christopher", "3"},
		{"Dan", "444"},
	}).Render()
}
```

</details>

### table/from-csv

![Animation](from-csv/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"encoding/csv"
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	data := `Name,Role,Commits
Ada,maintainer,1204
Grace,reviewer,388
Linus,contributor,57`

	// WithCSVReader reads every record. The CSV header line becomes the
	// table header.
	pterm.Table.WithCSVReader(csv.NewReader(strings.NewReader(data))).Render()
}
```

</details>

### table/from-structs

![Animation](from-structs/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type cache struct {
	Name  string
	State string
	Size  string
	Files int
	Path  string `pterm:"Location"`
	inode uint64
}

func main() {
	caches := []cache{
		{Name: "npm", State: "stale", Size: "1.20 GB", Files: 48211, Path: "~/.npm/_cacache"},
		{Name: "go", State: "fresh", Size: "2.84 GB", Files: 112530, Path: "~/go/pkg/mod"},
		{Name: "pip", State: "stale", Size: "87 MB", Files: 1302, Path: "~/.cache/pip"},
	}

	// Field names become the header. A pterm tag renames a column, and
	// unexported fields are left out.
	_ = pterm.Table.WithData(putils.TableFromStructs(caches)).Render()
}
```

</details>

### table/multiple-lines

![Animation](multiple-lines/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Cells can hold several lines. A row grows to fit its tallest cell, and
	// row separators keep the rows apart.
	pterm.Table.WithRowSeparators().WithData(pterm.TableData{
		{"Command", "Description"},
		{"init", "Create a new project\nin the current folder"},
		{"build", "Compile the project"},
		{"deploy", "Upload the build\nand restart the service\non every host"},
	}).Render()
}
```

</details>

### table/row-separators

![Animation](row-separators/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Row separators work in boxed tables too, where they join the frame.
	pterm.Table.WithBoxed().WithRowSeparators().WithData(pterm.TableData{
		{"Region", "Nodes", "Load"},
		{"eu-central", "12", "64%"},
		{"us-east", "30", "81%"},
		{"ap-south", "6", "23%"},
	}).Render()
}
```

</details>

### table/styles

![Animation](styles/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Styles use theme tones, so the table follows the active theme. Here the
	// status column is toned and the failed row stands out.
	pterm.Table.
		WithData(pterm.TableData{
			{"Job", "Status", "Duration"},
			{"lint", "passed", "12.4s"},
			{"test", "passed", "1m 3s"},
			{"build", "failed", "48.0s"},
			{"deploy", "skipped", "-"},
		}).
		WithColumnStyle(1, pterm.ToneAccent.Style()).
		WithRowStyle(2, pterm.ToneError.Style()).
		Render()

	pterm.Println()

	// An alternate row style makes long rows easier to follow.
	pterm.Table.
		WithData(pterm.TableData{
			{"Firstname", "Lastname", "Email"},
			{"Paul", "Dean", "paul@example.com"},
			{"Callie", "Mckay", "callie@example.com"},
			{"Libby", "Camacho", "libby@example.com"},
			{"张", "小宝", "zhang@example.com"},
		}).
		WithAlternateRowStyle(pterm.ToneMuted.Style()).
		Render()
}
```

</details>

