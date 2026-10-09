# spinner/multiple

![Animation](https://vhs.charm.sh/vhs-7DrAMyia1pQbonfygvudY9.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A MultiPrinter lets several spinners run at the same time. Each spinner
	// writes to its own writer obtained from the multi printer.
	multi := pterm.DefaultMultiPrinter

	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")
	spinner3, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 3")

	multi.Start()

	// Each spinner can resolve on its own while the others keep spinning.
	time.Sleep(time.Millisecond * 1000)
	spinner1.Success("Spinner 1 is done!")

	time.Sleep(time.Millisecond * 750)
	spinner2.Fail("Spinner 2 failed!")

	time.Sleep(time.Millisecond * 500)
	spinner3.Warning("Spinner 3 has a warning!")

	multi.Stop()
}
```
