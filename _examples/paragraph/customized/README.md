# paragraph/customized

![Animation](https://vhs.charm.sh/vhs-5DTRfS5OXeOv4yybnL8aOE.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	longText := "This is a custom paragraph printer. As you can see, no words are separated, " +
		"but the text is split at the spaces. This is useful for continuous text of all kinds. You can manually change the line width if you want to." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// The paragraph printer wraps at word boundaries. WithMaxWidth caps the
	// line width instead of using the full terminal width.
	pterm.DefaultParagraph.WithMaxWidth(60).Println(longText)

	pterm.Println()

	longTextWithoutParagraph := "This text is written with the default Println() function. No intelligent splitting here." +
		"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam"

	// For comparison: plain Println lets the terminal break lines wherever
	// they happen to overflow, even in the middle of a word.
	pterm.Println(longTextWithoutParagraph)
}
```
