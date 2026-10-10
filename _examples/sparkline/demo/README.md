# sparkline/demo

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// One column per value, scaled between the smallest and the largest.
	pterm.Sparkline.
		WithValues(3, 5, 4, 8, 12, 9, 14, 18, 16, 21, 19, 24, 22, 17, 13, 15).
		WithLabel("requests per second").
		Render()
}
```
