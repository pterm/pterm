### spinner/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create and start a fork of the default spinner.
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Some informational action...")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerInfo.Info()          // Resolve spinner with information message.

	// Create and start a fork of the default spinner.
	spinnerSuccess, _ := pterm.DefaultSpinner.Start("Doing something important... (will succeed)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerSuccess.Success()    // Resolve spinner with success message.

	// Create and start a fork of the default spinner.
	spinnerWarning, _ := pterm.DefaultSpinner.Start("Doing something important... (will warn)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerWarning.Warning()    // Resolve spinner with warning message.

	// Create and start a fork of the default spinner.
	spinnerFail, _ := pterm.DefaultSpinner.Start("Doing something important... (will fail)")
	time.Sleep(time.Second * 2) // Simulate 3 seconds of processing something.
	spinnerFail.Fail()          // Resolve spinner with error message.

	// Create and start a fork of the default spinner.
	spinnerNochange, _ := pterm.DefaultSpinner.Start("Checking something important... (will result in no change)")
	// Replace the InfoPrinter with a custom "NOCHG" one
	spinnerNochange.InfoPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.Style{pterm.FgLightBlue},
		Prefix: pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgLightBlue},
			Text:  " NOCHG ",
		},
	}
	time.Sleep(time.Second * 2)                     // Simulate 3 seconds of processing something.
	spinnerNochange.Info("No change were required") // Resolve spinner with error message.

	// Create and start a fork of the default spinner.
	spinnerLiveText, _ := pterm.DefaultSpinner.Start("Doing a lot of stuff...")
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.UpdateText("It's really much")   // Update spinner text.
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.UpdateText("We're nearly done!") // Update spinner text.
	time.Sleep(time.Second)                          // Simulate 2 seconds of processing something.
	spinnerLiveText.Success("Finally!")              // Resolve spinner with success message.
}

```

</details>

### spinner/multiple

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/multiple/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer. This allows multiple spinners to print simultaneously.
	multi := pterm.DefaultMultiPrinter

	// Create and start spinner 1 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 1".
	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")

	// Create and start spinner 2 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 2".
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	// Create and start spinner 3 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 3".
	spinner3, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 3")

	// Start the multi printer. This will start printing all the spinners.
	multi.Start()

	// Wait for 1 second.
	time.Sleep(time.Millisecond * 1000)

	// Stop spinner 1 with a success message.
	spinner1.Success("Spinner 1 is done!")

	// Wait for 750 milliseconds.
	time.Sleep(time.Millisecond * 750)

	// Stop spinner 2 with a failure message.
	spinner2.Fail("Spinner 2 failed!")

	// Wait for 500 milliseconds.
	time.Sleep(time.Millisecond * 500)

	// Stop spinner 3 with a warning message.
	spinner3.Warning("Spinner 3 has a warning!")

	// Stop the multi printer. This will stop printing all the spinners.
	multi.Stop()
}

```

</details>

