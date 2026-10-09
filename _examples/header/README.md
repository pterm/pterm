### header/demo

![Animation](https://vhs.charm.sh/vhs-3GdCjND3XJAaMx79ktH51Q.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### header/custom

![Animation](https://vhs.charm.sh/vhs-5c8QzHYrjymhxB5DTG6XEF.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Restyle the default header on the fly. The margin adds horizontal
	// padding on both sides of the text.
	pterm.DefaultHeader.WithMargin(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("This is a custom header!")

	// Alternatively, build a HeaderPrinter from scratch instead of deriving
	// from DefaultHeader.
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}

	newHeader.Println("This is a custom header!")
}
```

</details>

