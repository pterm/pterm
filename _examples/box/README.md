### box/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A box frames text that belongs together. The title sits in the
	// border, so the frame costs no extra line.
	pterm.Box.WithTitle("Result").Println("3 files changed\n128 insertions, 41 deletions")

	// Boxes are text printers, so their output can go into other printers,
	// even into another box.
	inner := pterm.Box.WithTitle("Cache").Sprint("12 entries\n4.1 MB")
	pterm.Box.WithTitle("Status").Println(inner)
}
```

</details>

### box/padding

![Animation](padding/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// By default there is one column of space left and right of the text.
	pterm.Box.Println("Default padding")

	// Padding can be set for all sides at once or for each side.
	pterm.Box.WithPadding(1).Println("One line and column on every side")
	pterm.Box.WithHorizontalPadding(4).Println("Wide sides")
	pterm.Box.WithLeftPadding(0).WithRightPadding(0).Println("No padding")
}
```

</details>

### box/title-position

![Animation](title-position/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The title can sit on the left, in the center or on the right, at the
	// top or at the bottom of the frame.
	pterm.Box.WithTitle("top left").WithWidth(40).Println("The frame keeps its width.")
	pterm.Box.WithTitle("top center").WithTitlePosition(pterm.BoxTitleTopCenter).WithWidth(40).Println("The title moves.")
	pterm.Box.WithTitle("bottom right").WithTitlePosition(pterm.BoxTitleBottomRight).WithWidth(40).Println("Even to the bottom.")
}
```

</details>

