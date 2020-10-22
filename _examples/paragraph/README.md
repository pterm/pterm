# paragraph

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.DefaultParagraph.Println("This is the default paragraph printer. As you can see, no words are separated, " +
		"but the letters is split at the spaces. This is useful for continuous letters of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")

	pterm.Println()

	pterm.Println("This letters is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam")
}

```
