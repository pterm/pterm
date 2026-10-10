# box/padding

![Animation](animation.gif)

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
