package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// spinnerSuccess := pterm.DefaultSpinner.Start("Doing something important... (will succeed)")

	// time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	// spinnerSuccess.Success()

	// spinnerWarning := pterm.DefaultSpinner.Start("Doing something important... (will warn)")

	// time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	// spinnerWarning.Warning()

	// spinnerFail := pterm.DefaultSpinner.Start("Doing something important... (will fail)")

	// time.Sleep(time.Second * 3) // Simulate 3 seconds of processing something

	// spinnerFail.Fail()

	spinnerLiveText := pterm.DefaultSpinner.Start("Doing a lot of stuff...")

	time.Sleep(time.Second * 2)

	spinnerLiveText.UpdateText("It's really much")

	time.Sleep(time.Second * 2)

	spinnerLiveText.UpdateText("We're nearly done!")

	time.Sleep(time.Second * 2)

	spinnerLiveText.Success("Finally!")

	spinnerMessageText := pterm.DefaultSpinner.Start("🤔 Searching for something tacky to say ...")

	time.Sleep(time.Second * 2)

	spinnerMessageText.UpdateText(" 🤷 Lots of stuff out there...")

	time.Sleep(time.Second * 2)

	spinnerMessageText.UpdateText("Here it is")

	time.Sleep(time.Second * 1)

	motivation := 100

	spinnerMessageText.PassMessage("Don’t wait for someone else to do it\nHire yourself and start calling the shots\nand keep it %d 👊", motivation)
}
