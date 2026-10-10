# spinner/demo

![Animation](animation.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	spinner, _ := pterm.Spinner.Start("Resolving dependencies")
	time.Sleep(2 * time.Second)
	spinner.UpdateText("Downloading 42 packages")
	time.Sleep(2 * time.Second)

	// Output printed while the spinner runs appears above it.
	pterm.Info.Println("Using the module cache at ~/go/pkg/mod")
	time.Sleep(time.Second)

	spinner.Success("Installed 42 packages")
}
```
