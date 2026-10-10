# paragraph/demo

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A paragraph wraps running text at word boundaries. Lines stay at most
	// 80 columns wide, which keeps text easy to read on wide terminals.
	// A blank line starts a new paragraph.
	pterm.Paragraph.Println(`pterm prints beautiful terminal output with one line of code. Every printer
follows one theme, so tables, spinners and prompts look like parts of one tool.

Output adapts to the terminal: colors are turned off when the output is a pipe,
and text wraps to the width of the window.`)
}
```
