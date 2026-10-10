# keyvalue/demo

![Animation](animation.gif)

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
