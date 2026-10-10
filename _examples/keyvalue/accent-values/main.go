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
