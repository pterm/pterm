# box/title-position

![Animation](animation.gif)

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
