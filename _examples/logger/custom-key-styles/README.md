# logger/custom-key-styles

![Animation](https://vhs.charm.sh/vhs-XImFCwmCtfj6gxSAL1nPL.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// WithKeyStyles replaces the whole key style map, so only the keys listed
	// here get a custom style.
	logger = logger.WithKeyStyles(map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	})

	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// AppendKeyStyle adds a single key style on top of the existing ones.
	logger.AppendKeyStyle("foo", *pterm.NewStyle(pterm.FgBlue))

	logger.Info("The foo key should now be blue", logger.Args("priority", "low", "foo", "bar"))
}
```
