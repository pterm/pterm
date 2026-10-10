### keyvalue/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Keys line up and stay muted, so the values are what the eye lands on.
	// A suffix adds a muted note after a value, and long values wrap under
	// themselves instead of under the keys.
	pterm.KeyValue.WithPairs(
		pterm.KV("Version", "0.1.0"),
		pterm.KV("Path", "~/.local/fleet/dotfiles/src/cli/main.ts"),
		pterm.KV("Uptime", "1m 34s").WithSuffix("since last boot"),
		pterm.KV("Commands", 12),
		pterm.KV("Description", "A small command line tool that keeps the dotfiles of every machine in the fleet in sync and reports what drifted."),
	).Render()
}
```

</details>

### keyvalue/accent-values

![Animation](accent-values/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Accent values suit output where the values are versions, paths and
	// numbers worth noticing. Indent 0 starts the keys at the left edge.
	pterm.KeyValue.WithAccentValues().WithIndent(0).WithPairs(
		pterm.KV("Go", "go1.27.2"),
		pterm.KV("Module", "github.com/pterm/pterm"),
		pterm.KV("Binary", "12.4").WithSuffix("MB"),
		pterm.KV("Build", "3.2s"),
	).Render()
}
```

</details>

### keyvalue/from-struct

![Animation](from-struct/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type build struct {
	Version  string
	Commit   string
	Target   string `pterm:"Target platform"`
	Duration time.Duration
	Cached   bool
}

func main() {
	b := build{
		Version:  "v0.13.0",
		Commit:   "4ce46dd",
		Target:   "linux/amd64",
		Duration: 4200 * time.Millisecond,
		Cached:   true,
	}

	// Every exported field becomes a pair, in the order it is declared.
	_ = pterm.KeyValue.WithPairs(putils.KeyValueFromStruct(b)...).Render()
}
```

</details>

