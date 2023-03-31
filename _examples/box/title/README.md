# box/title

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Default titled bpx
	paddedBox := pterm.DefaultBox.WithLeftPadding(4).WithRightPadding(4).WithTopPadding(1).WithBottomPadding(1)

	title := pterm.LightRed("I'm a box!")

	box1 := paddedBox.WithTitle(title).Sprint("Hello, World!\n      1")
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
