# spinner/statuses

![Animation](animation.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A spinner ends as a status line. Pick the ending that matches what
	// happened.
	s, _ := pterm.Spinner.Start("Compiling")
	time.Sleep(time.Second)
	s.Success("Compiled 18 packages")

	s, _ = pterm.Spinner.Start("Running tests")
	time.Sleep(time.Second)
	s.Warning("2 tests were skipped")

	s, _ = pterm.Spinner.Start("Uploading coverage")
	time.Sleep(time.Second)
	s.Fail("Coverage upload timed out")

	s, _ = pterm.Spinner.Start("Publishing docs")
	s.Skip("Docs are unchanged")
}
```
