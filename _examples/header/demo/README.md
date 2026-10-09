# header/demo

![Animation](https://vhs.charm.sh/vhs-3GdCjND3XJAaMx79ktH51Q.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// By default the header is only as wide as its content plus the margin.
	pterm.DefaultHeader.Println("This is the default header!")

	pterm.Println()

	// WithFullWidth stretches the header background across the whole terminal.
	pterm.DefaultHeader.WithFullWidth().Println("This is a full-width header.")
}
```
