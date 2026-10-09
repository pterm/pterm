### bigtext/demo

![Animation](https://vhs.charm.sh/vhs-Dltv80znAaeuvbEhngzLB.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Big ASCII-art text in the default theme style. Handy for title screens.
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()

	// Each letter group can carry its own style, so parts of the text can be
	// colored independently.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle()),
	).Render()

	// TrueColor works too. On terminals without TrueColor support, PTerm
	// downsamples the RGB value automatically.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithRGB("PTerm", pterm.NewRGB(255, 215, 0)),
	).Render()
}
```

</details>

### bigtext/colored

![Animation](https://vhs.charm.sh/vhs-6OVbbx3QPZS9Rz75SwAzmR.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// Each LettersFromStringWithStyle call gets its own color, so parts of
	// the big text can be highlighted independently.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()
}
```

</details>

### bigtext/default

![Animation](https://vhs.charm.sh/vhs-3AzGuVlmJ9yL1ADI8F1IV2.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	// BigText takes Letters rather than a plain string; putils converts one
	// into the other.
	letters := putils.LettersFromString("PTerm")

	pterm.DefaultBigText.WithLetters(letters).Render()
}
```

</details>

