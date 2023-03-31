### box/custom-padding

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/custom-padding/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println("Hello, World!")
}

```

</details>

### box/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultBox.Println("Hello, World!")
}

```

</details>

### box/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Info.Println("This might not be rendered correctly on GitHub,\nbut it will work in a real terminal.\nThis is because GitHub does not use a monospaced font by default for SVGs")

	panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt\nut labore et dolore\nmagna aliqua.")
	panel2 := pterm.DefaultBox.WithTitle("title").Sprint("Ut enim ad minim veniam,\nquis nostrud exercitation\nullamco laboris\nnisi ut aliquip\nex ea commodo\nconsequat.")
	panel3 := pterm.DefaultBox.WithTitle("bottom center title").WithTitleBottomCenter().Sprint("Duis aute irure\ndolor in reprehenderit\nin voluptate velit esse cillum\ndolore eu fugiat\nnulla pariatur.")

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: panel1}, {Data: panel2}},
		{{Data: panel3}},
	}).Srender()

	pterm.DefaultBox.WithTitle("Lorem Ipsum").WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(panels)
}

```

</details>

### box/title

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/title/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

