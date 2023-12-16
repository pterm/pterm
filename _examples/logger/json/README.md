# logger/json

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with Trace level and JSON formatter
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithFormatter(pterm.LogFormatterJSON)

	// Log a Trace level message with additional arguments
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff
	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a Debug level message with arguments from the map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	// Log Info, Warn, Error, and Fatal level messages with additional arguments
	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```
