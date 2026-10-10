# header/without-badge

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Without badge text, the title itself goes into the badge.
	pterm.Header.Println("Release notes")

	// The rule can match the width of the header line instead of the
	// terminal, and the badge can use any tone.
	pterm.Header.
		WithBadge("beta").
		WithBadgeTone(pterm.ToneWarning).
		WithFullWidth(false).
		Println("Features that may still change")
}
```
