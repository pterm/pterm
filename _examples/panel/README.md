### panel/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Boxed panels make a small dashboard. Every row is laid out side by side
	// and frames of one row end on the same line.
	pterm.Panel.WithBoxed().
		WithRow(
			pterm.PanelItem{Title: "Uptime", Data: "99.99%\nlast 30 days"},
			pterm.PanelItem{Title: "Requests", Data: "1.2M per day"},
			pterm.PanelItem{Title: "Errors", Data: "0.02%\n14 in the last hour"},
		).
		WithRow(
			pterm.PanelItem{Title: "Region", Data: "eu-central-1"},
			pterm.PanelItem{Title: "Version", Data: "v2.4.1\nbuilt 5 minutes ago"},
		).
		Render()
}
```

</details>

### panel/columns

![Animation](columns/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Without frames, panels are plain columns. Their content can come from
	// other printers, such as these bullet lists.
	tools, _ := pterm.BulletList.WithTexts("npm", "bun", "pip").Srender()
	stores, _ := pterm.BulletList.WithTexts("Docker images", "Go modules").Srender()

	pterm.Panel.WithPadding(6).WithPanels(pterm.Panels{{
		{Title: "Package managers", Data: tools},
		{Title: "Stores", Data: stores},
	}}).Render()
}
```

</details>

### panel/same-column-width

![Animation](same-column-width/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// SameColumnWidth widens every panel to the widest one of its column,
	// so the frames of different rows line up.
	pterm.Panel.WithBoxed().WithSameColumnWidth().WithPanels(pterm.Panels{
		{
			{Title: "CPU", Data: "12%"},
			{Title: "Memory", Data: "3.1 GB of 16 GB"},
		},
		{
			{Title: "Disk", Data: "212 GB free"},
			{Title: "Network", Data: "4 MB/s"},
		},
	}).Render()
}
```

</details>

