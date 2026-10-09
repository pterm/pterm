# spinner/demo

![Animation](https://vhs.charm.sh/vhs-3Zy8BjZTjS9OkbJkLLflis.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A spinner can resolve as Info, Success, Warning or Fail. The spinner
	// line is replaced by the matching prefix printer output.
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Some informational action...")
	time.Sleep(time.Second * 2)
	spinnerInfo.Info()

	spinnerSuccess, _ := pterm.DefaultSpinner.Start("Doing something important... (will succeed)")
	time.Sleep(time.Second * 2)
	spinnerSuccess.Success()

	spinnerWarning, _ := pterm.DefaultSpinner.Start("Doing something important... (will warn)")
	time.Sleep(time.Second * 2)
	spinnerWarning.Warning()

	spinnerFail, _ := pterm.DefaultSpinner.Start("Doing something important... (will fail)")
	time.Sleep(time.Second * 2)
	spinnerFail.Fail()

	// The resolve printers are plain PrefixPrinters, so they can be swapped
	// out. Here Info resolves with a custom "NOCHG" prefix instead.
	spinnerNochange, _ := pterm.DefaultSpinner.Start("Checking something important... (will result in no change)")
	spinnerNochange.InfoPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.Style{pterm.FgLightBlue},
		Prefix: pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgLightBlue},
			Text:  " NOCHG ",
		},
	}

	time.Sleep(time.Second * 2)
	spinnerNochange.Info("No changes were required")

	// The text can be updated while the spinner keeps running.
	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Doing a lot of stuff...")
	time.Sleep(time.Second)
	spinnerLiveText.UpdateText("It's really much")
	time.Sleep(time.Second)
	spinnerLiveText.UpdateText("We're nearly done!")
	time.Sleep(time.Second)
	spinnerLiveText.Success("Finally!")
}
```
