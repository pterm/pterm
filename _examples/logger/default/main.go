package main

import "github.com/pterm/pterm"

func main() {
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace)

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	logger.Fatal("Oh no, we are doomed!", logger.Args("fatal", true))
}
