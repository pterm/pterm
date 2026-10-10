### rule/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A rule separates parts of the output without the weight of a section
	// title. It spans the terminal.
	pterm.Text.Println("Build finished in 4.2s")
	pterm.Rule.Println()

	// The text becomes a title inside the rule.
	pterm.Rule.Println("Test results")
	pterm.Text.Println("42 passed, 0 failed")
}
```

</details>

### rule/width

![Animation](width/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A fixed width lines the rule up with other output.
	pterm.Rule.WithWidth(40).Println("Summary")

	// The title can be muted, and the line can use another glyph and tone.
	pterm.Rule.
		WithWidth(40).
		WithTitleStyle(pterm.ToneMuted.Style()).
		WithGlyph("=").
		WithTone(pterm.ToneMuted).
		Println("Details")
}
```

</details>

