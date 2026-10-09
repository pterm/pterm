# paragraph/demo

![Animation](https://vhs.charm.sh/vhs-rDsuWNMNIZlEJ4RMSFwhR.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The paragraph printer wraps long text at word boundaries so it fits the
	// terminal width without breaking words apart.
	pterm.DefaultParagraph.Println("This is the default paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")

	pterm.Println()

	// For comparison: plain Println lets the terminal break lines wherever
	// they happen to overflow, even in the middle of a word.
	pterm.Println("This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")
}
```
