# logger/with-caller

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Create a logger with Trace level and caller information
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithCaller()

	// Log a trace message with additional arguments
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff
	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug message with arguments from a map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	// Log an info message with additional arguments
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Log a warning message with additional arguments
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Log an error message with additional arguments
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Log an info message with additional arguments. PTerm will automatically wrap long logs.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Log a fatal message with additional arguments. This will terminate the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```
