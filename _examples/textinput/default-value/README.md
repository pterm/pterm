# textinput/default-value

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The default value is shown muted until the user types. Enter on an
	// empty input returns it, and tab copies it into the input to edit it.
	branch, _ := pterm.TextInput.
		WithDefaultValue("main").
		Show("Branch to deploy")

	pterm.Info.Printfln("Deploying %s", branch)
}
```
