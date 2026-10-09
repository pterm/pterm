### box/demo

![Animation](https://vhs.charm.sh/vhs-1YpsH81JEuvjgD6uxQ5iv3.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Boxes render to strings via Sprint, so they can be nested inside other
	// printers. Titles can be placed on any side of the border.
	panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	panel3 := pterm.DefaultBox.WithTitle("bottom center title").WithTitleBottomCenter().Sprint("Duis aute irure\ndolor in reprehenderit\nin voluptate velit esse cillum\ndolore eu fugiat\nnulla pariatur.")

	// Arrange the boxes in a grid: one row with two panels, one row with one.
	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: panel1}, {Data: panel2}},
		{{Data: panel3}},
	}).Srender()

	// Wrap the whole grid in an outer box.
	pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
}
```

</details>

### box/custom-padding

![Animation](https://vhs.charm.sh/vhs-cy3y8ESzFiFughcKeRmzu.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Padding adds empty space between the box border and its content,
	// configurable per side.
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Hello, World!")
}
```

</details>

### box/default

![Animation](https://vhs.charm.sh/vhs-1DIyAbs7eXHIzVW2Z7p4Md.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The box sizes itself to fit whatever it prints.
	pterm.DefaultBox.Println("Hello, World!")
}
```

</details>

### box/title

![Animation](https://vhs.charm.sh/vhs-3xc4Z0HU2WGZnIaPCUjymS.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// One box per title position. WithTitle* methods return a modified copy,
	// so paddedBox can be reused as a template without being changed.
	paddedBox := pterm.DefaultBox.WithLeftPadding(4).WithRightPadding(4).WithTopPadding(1).WithBottomPadding(1)

	// Titles may contain styled text.
	title := pterm.LightRed("I'm a box!")

	box1 := paddedBox.WithTitle(title).Sprint("Hello, World!\n      1") // top left is the default
	box2 := paddedBox.WithTitle(title).WithTitleTopCenter().Sprint("Hello, World!\n      2")
	box3 := paddedBox.WithTitle(title).WithTitleTopRight().Sprint("Hello, World!\n      3")
	box4 := paddedBox.WithTitle(title).WithTitleBottomRight().Sprint("Hello, World!\n      4")
	box5 := paddedBox.WithTitle(title).WithTitleBottomCenter().Sprint("Hello, World!\n      5")
	box6 := paddedBox.WithTitle(title).WithTitleBottomLeft().Sprint("Hello, World!\n      6")
	box7 := paddedBox.WithTitle(title).WithTitleTopLeft().Sprint("Hello, World!\n      7")

	pterm.DefaultPanel.WithPanels([][]pterm.Panel{
		{{box1}, {box2}, {box3}},
		{{box4}, {box5}, {box6}},
		{{box7}},
	}).Render()
}
```

</details>

