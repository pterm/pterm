### logger/custom-key-styles

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/logger/custom-key-styles/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace) // Only show logs with a level of Trace or higher.

	// Overwrite all key styles with a new map
	logger = logger.WithKeyStyles(map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	})

	// The priority key should now be red
	logger.Info("The priority key should now be red", logger.Args("priority", "low", "foo", "bar"))

	// Append a key style to the exisiting ones
	logger.AppendKeyStyle("foo", *pterm.NewStyle(pterm.FgBlue))

	// The foo key should now be blue
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

</details>

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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace)

	logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

	sleep()

	interstingStuff := map[string]any{
		"when were crayons invented":  "1903",
		"what is the meaning of life": 42,
		"is this interesting":         true,
	}
	logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
	sleep()

	logger.Info("That was actually interesting", logger.Args("such", "wow"))
	sleep()
	logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
	sleep()
	logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
	sleep()
	logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	sleep()
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

func sleep() {
	time.Sleep(time.Second * 3)
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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace).       // Only show logs with a level of Trace or higher.
		WithFormatter(pterm.LogFormatterJSON) // ! Make the logger print JSON logs.

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
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace). // Only show logs with a level of Trace or higher.
		WithCaller()                    // ! Show the caller of the log function.

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
	logger.Fatal("Oh no, this process is getting killed!", logger.Args("fatal", true))
}

```

</details>

