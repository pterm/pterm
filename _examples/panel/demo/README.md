# panel/demo

![Animation](https://vhs.charm.sh/vhs-1ojPuB9SETuUW6zO274YLY.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Panels form a 2D grid: the outer slice holds rows, the inner slices hold
	// the panels of each row. Panel content can be multiline and may come from
	// other printers.
	panels := pterm.Panels{
		{
			{Data: "This is the first panel"},
			{Data: pterm.DefaultHeader.Sprint("Hello, World!")},
			{Data: "This\npanel\ncontains\nmultiple\nlines"},
		},
		{
			{Data: pterm.Red("This is another\npanel line")},
			{Data: "This is the second panel\nwith a new line"},
		},
	}

	// Padding controls the horizontal space between panels in a row.
	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}
```
