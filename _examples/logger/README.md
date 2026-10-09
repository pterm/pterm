### logger/demo

![Animation](https://vhs.charm.sh/vhs-5ZEL81G8v6BoyEteQgEeTI.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>

### logger/custom-key-styles

![Animation](https://vhs.charm.sh/vhs-XImFCwmCtfj6gxSAL1nPL.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### logger/default

![Animation](https://vhs.charm.sh/vhs-1L9BFUa3Jk9MMxZUxBowEA.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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

	time.Sleep(time.Second * 2)

	// Fatal logs the message and then exits the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}
```

</details>

### logger/json

![Animation](https://vhs.charm.sh/vhs-4ZkTNA9mnIREsQGFxbg6iH.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The JSON formatter emits one JSON object per log line, which is handy
	// for machine-readable output in production.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithFormatter(pterm.LogFormatterJSON)

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
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Fatal logs the message and then exits the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}
```

</details>

### logger/with-caller

![Animation](https://vhs.charm.sh/vhs-7nMQzoh4SQfQBmEDun2JHP.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

