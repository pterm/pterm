# logger/demo

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	// Create a logger with trace level
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Log a trace level message
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Pause for 3 seconds
	sleep()

	// Define a map with interesting stuff
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug level message with arguments from the map
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	// Pause for 3 seconds
	sleep()

	// Log an info level message
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Pause for 3 seconds
	sleep()

	// Log a warning level message
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Pause for 3 seconds
	sleep()

	// Log an error level message
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Pause for 3 seconds
	sleep()

	// Log an info level message with a long text that will be automatically wrapped
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Pause for 3 seconds
	sleep()

	// Log a fatal level message
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

// Function to pause the execution for 3 seconds
func sleep() {
	time.Sleep(time.Second * 3)
}

```
