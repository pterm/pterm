### text/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Text prints like fmt, but in the theme's text color, and long lines
	// wrap at word boundaries to fit the terminal.
	pterm.Text.Println("Plain text follows the theme, so it matches every other printer.")
	pterm.Text.Println("Long lines wrap at the terminal width instead of being cut in the middle of a word, which keeps help texts and long messages readable in narrow windows.")
	pterm.Text.Printfln("Formatting works too: %d files, %s.", 3, "1.2 MB")
}
```

</details>

### text/styled

![Animation](styled/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Any tone of the theme can be used. Keep tones for meaning: muted
	// text for hints, accent for values worth noticing.
	pterm.Text.WithTone(pterm.ToneMuted).Println("Hints and metadata are muted.")
	pterm.Text.WithTone(pterm.ToneAccent).Println("Values stand out in the accent color.")

	// WithNoWrap keeps lines exactly as they are, for preformatted output.
	pterm.Text.WithNoWrap().Println("id    name     status\n42    worker   running")
}
```

</details>

