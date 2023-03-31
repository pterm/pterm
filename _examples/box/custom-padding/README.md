# box/custom-padding

![Animation](animation.svg)

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
