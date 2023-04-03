package main

import "github.com/pterm/pterm"

func main() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace) // Only show logs with a level of Trace or higher.

	// Overwrite all key styles with a new map
	logger = logger.WithKeyStyles(map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	})

	// The priority key should now be red
	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// Append a key style to the exisiting ones
	logger.AppendKeyStyle("foo", *pterm.NewStyle(pterm.FgBlue))

	// The foo key should now be blue
	logger.Info("The foo key should now be blue", logger.Args("priority", "low", "foo", "bar"))
}
