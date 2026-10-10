# center/each-line

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// For loose text, every line can be centered on its own.
	pterm.Center.WithCenterEachLineSeparately().Println("Thanks for trying pterm\nStar it on GitHub if you like it\nhttps://github.com/pterm/pterm")
}
```
