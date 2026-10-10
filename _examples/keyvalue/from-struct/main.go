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
