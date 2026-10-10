### bigtext/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Big letters for title screens, in a gradient from Primary to Accent.
	pterm.BigText.Println("PTerm")
}
```

</details>

### bigtext/gradient

![Animation](gradient/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A gradient can run through any number of tones.
	pterm.BigText.WithGradient(pterm.ToneInfo, pterm.TonePrimary, pterm.ToneError).Println("v1.0")

	pterm.Println()

	// The offset shifts the gradient by a share of the text width. It
	// bounces at the ends, so stepping it from 0 to 2 in a loop sweeps the
	// colors across the text and back without a jump.
	pterm.BigText.WithGradientOffset(0.5).Println("PTerm")
}
```

</details>

### bigtext/letters

![Animation](letters/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Letters with a style of their own keep it; the rest follow the
	// printer's gradient.
	pterm.BigText.WithLetters(
		pterm.LettersFromString("P", pterm.ToneAccent.Style()),
		pterm.LettersFromString("Term"),
	).Render()
}
```

</details>

### bigtext/tone

![Animation](tone/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// One tone instead of the gradient.
	pterm.BigText.WithTone(pterm.ToneSuccess).Println("Done")
}
```

</details>

