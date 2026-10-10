# panel/demo

![Animation](animation.gif)

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
