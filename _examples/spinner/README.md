### spinner/demo

![Animation](https://vhs.charm.sh/vhs-3Zy8BjZTjS9OkbJkLLflis.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

### spinner/custom

![Animation](https://vhs.charm.sh/vhs-7yMMUFKhhlDSkKnbXrdVEO.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The spinner animation is just a sequence of frames, so any set of
	// strings works. WithStyle colors the animation.
	spinner, _ := pterm.DefaultSpinner.
		WithSequence("▁", "▃", "▅", "▇", "▅", "▃").
		WithStyle(pterm.NewStyle(pterm.FgCyan)).
		Start("Uploading assets...")

	time.Sleep(time.Second * 2)

	// The text can change while the spinner keeps running.
	spinner.UpdateText("Finalizing upload...")
	time.Sleep(time.Second * 2)

	spinner.Success("Upload complete")
}
```

</details>

### spinner/multiple

![Animation](https://vhs.charm.sh/vhs-7DrAMyia1pQbonfygvudY9.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

