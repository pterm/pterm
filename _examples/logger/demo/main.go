package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The default log level is Info. Lower it to Trace so every level shows up.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// logger.Args pairs up keys and values for structured output.
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	sleep()

	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// ArgsFromMap turns an existing map into logger arguments.
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	sleep()

	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	sleep()

	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	sleep()

	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	sleep()

	// Long messages are wrapped to the terminal width automatically.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	sleep()

	// Fatal logs the message and then exits the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

func sleep() {
	time.Sleep(time.Second * 3)
}
