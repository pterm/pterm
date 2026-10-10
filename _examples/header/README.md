### header/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A header opens the output of a program: its name in a badge, what is
	// running and a short muted description, with a rule underneath.
	pterm.Header.
		WithBadge("APP").
		WithSubtitle("Design system gallery").
		Println("debug ui")

	pterm.Text.Println("Everything below the header belongs to this run.")
}
```

</details>

### header/without-badge

![Animation](without-badge/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Without badge text, the title itself goes into the badge.
	pterm.Header.Println("Release notes")

	// The rule can match the width of the header line instead of the
	// terminal, and the badge can use any tone.
	pterm.Header.
		WithBadge("beta").
		WithBadgeTone(pterm.ToneWarning).
		WithFullWidth(false).
		Println("Features that may still change")
}
```

</details>

