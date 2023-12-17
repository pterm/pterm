package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with a level of Trace or higher.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Define a new style for the "priority" key.
	priorityStyle := map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	}

	// Overwrite all key styles with the new map.
	logger = logger.WithKeyStyles(priorityStyle)

	// Log an info message. The "priority" key will be displayed in red.
	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// Define a new style for the "foo" key.
	fooStyle := *pterm.NewStyle(pterm.FgBlue)

	// Append the new style to the existing ones.
	logger.AppendKeyStyle("foo", fooStyle)

	// Log another info message. The "foo" key will be displayed in blue.
	logger.Info("The foo key should now be blue", logger.Args("priority", "low", "foo", "bar"))
}
