### callout/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Callout.Note("Note").Println("Neutral information that supports the command.")
	pterm.Println()
	pterm.Callout.Warn("Heads up").Println("Something to double check before continuing.")
	pterm.Println()
	pterm.Callout.Tip("Done").Println("This is the good outcome.")
	pterm.Println()
	pterm.Callout.Danger("Error").Println("Something went wrong while reading the config.")
}
```

</details>

### callout/hint

![Animation](hint/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A hint is a muted last line, good for the cause of an error or a
	// next step.
	pterm.Callout.Danger("Error").
		WithHint("caused by: ENOENT: no such file").
		Println("Something went wrong while reading the config\nRun `app debug doctor` to see where app is looking.")
}
```

</details>

