# box/custom-padding

![Animation](https://vhs.charm.sh/vhs-i8JKBbJhrtKmtrcIN2uCr.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a default box with custom padding options and print "Hello, World!" inside it.
	pterm.DefaultBox.WithRightPadding(10).WithLeftPadding(10).WithTopPadding(2).WithBottomPadding(2).Println("Hello, World!")
}
```
