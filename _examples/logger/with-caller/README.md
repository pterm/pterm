# logger/with-caller

![Animation](https://vhs.charm.sh/vhs-7nMQzoh4SQfQBmEDun2JHP.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// WithCaller adds the file and line of the log call to every message.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithCaller()

	// logger.Args pairs up keys and values for structured output.
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	interestingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// ArgsFromMap turns an existing map into logger arguments.
	logger.Debug("This might be interesting", logger.ArgsFromMap(interestingStuff))

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Long messages are wrapped to the terminal width automatically.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Fatal logs the message and then exits the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}
```
