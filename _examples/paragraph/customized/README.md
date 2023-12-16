# paragraph/customized

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define a long text to be printed as a paragraph.
	longText := "This is a custom paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// Print the long text as a paragraph with a custom maximal width of 60 characters.
	pterm.DefaultParagraph.WithMaxWidth(60).Println(longText)

	// Print a line space to separate the paragraph from the following text.
	pterm.Println()

	// Define another long text to be printed without a paragraph printer.
	longTextWithoutParagraph := "This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// Print the long text without using a paragraph printer.
	pterm.Println(longTextWithoutParagraph)
}

```
