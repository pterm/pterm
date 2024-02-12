### logger/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### logger/custom-key-styles

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/custom-key-styles/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
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

```

</details>

### logger/default

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/default/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	// Create a logger with a level of Trace or higher.
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	// Log a trace message with additional arguments.
	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	// Create a map of interesting stuff.
	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}

	// Log a debug message with arguments from a map.
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))

	// Log an info message with additional arguments.
	logger.Info("That was actually interesting", logger.Args("such", "wow"))

	// Log a warning message with additional arguments.
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))

	// Log an error message with additional arguments.
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))

	// Log an info message with additional arguments. PTerm will automatically wrap long logs.
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))

	// Pause for 2 seconds.
	time.Sleep(time.Second * 2)

	// Log a fatal message with additional arguments. This will terminate the process.
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```

</details>

### logger/json

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/json/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### logger/with-caller

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/with-caller/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

