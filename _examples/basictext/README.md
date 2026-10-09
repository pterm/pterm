### basictext/demo

![Animation](https://vhs.charm.sh/vhs-5eF1XRSmYJjH9IO6AcLXrR.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// DefaultBasicText prints plain, unstyled text. Its value is that it
	// satisfies the TextPrinter interface, so it can be passed anywhere a
	// styled printer would go.
	pterm.DefaultBasicText.Println("Default basic text printer.")
	pterm.DefaultBasicText.Println("Can be used in any" + pterm.LightMagenta(" TextPrinter ") + "context.")
}
```

</details>

