# theme/env

![Animation](animation.gif)

```go
package main

import (
	"github.com/pterm/pterm"
)

// Run this example with different environment variables to see the theme
// change without touching the code:
//
//	PTERM_THEME=latte go run .
//	PTERM_THEME=ansi go run .
//	PTERM_THEME_PRIMARY=#f5a97f PTERM_THEME_SUCCESS=cyan go run .
//
// PTERM_THEME defaults to "auto", which asks the terminal for its
// background color and picks Latte on light backgrounds.
func main() {
	pterm.Info.Println("Theme slots come from the environment")
	pterm.Success.Println("Tests passed")
	pterm.Warning.Println("Coverage dropped by 2%")
	pterm.Println(pterm.TonePrimary.Bold().Sprint("primary"), pterm.ToneAccent.Sprint("accent"), pterm.ToneMuted.Sprint("muted"))
}
```
