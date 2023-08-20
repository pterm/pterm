package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	multi := pterm.DefaultMultiPrinter

	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")
	spinner3, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 3")

	multi.Start()

	time.Sleep(time.Millisecond * 1000)
	spinner1.Success("Spinner 1 is done!")
	time.Sleep(time.Millisecond * 750)
	spinner2.Fail("Spinner 2 failed!")
	time.Sleep(time.Millisecond * 500)
	spinner3.Warning("Spinner 3 has a warning!")

	multi.Stop()
}
