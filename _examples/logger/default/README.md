# logger/default

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace) // Only show logs with a level of Trace or higher.

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// You can also use the `ArgsFromMap` function to create a `Args` object from a map.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	time.Sleep(time.Second * 2)
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```
