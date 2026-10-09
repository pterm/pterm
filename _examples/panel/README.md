### panel/demo

![Animation](https://vhs.charm.sh/vhs-1ojPuB9SETuUW6zO274YLY.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### panel/boxed

![Animation](https://vhs.charm.sh/vhs-3am3fCMbszVAZx96utD5Z3.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Panels form a 2D grid: the outer slice holds rows, the inner slices hold
	// the panels of each row.
	panels := pterm.Panels{
		{
			{Data: "Uptime\n99.99%"},
			{Data: "Requests\n1.2M/day"},
			{Data: "Errors\n0.02%"},
		},
		{
			{Data: "Region\neu-central"},
			{Data: "Version\nv2.4.1"},
			{Data: "Build\n#4821"},
		},
	}

	// WithBoxPrinter draws each panel inside a box. WithSameColumnWidth pads
	// every panel of a column to the widest one, so the boxes line up.
	_ = pterm.DefaultPanel.
		WithPanels(panels).
		WithBoxPrinter(pterm.DefaultBox).
		WithSameColumnWidth().
		Render()
}
```

</details>

