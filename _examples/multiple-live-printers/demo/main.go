package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The multi printer renders several live printers at once. Each live
	// printer gets its own line via multi.NewWriter().
	multi := pterm.DefaultMultiPrinter

	spinner1, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 1")
	spinner2, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Spinner 2")

	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	// Nothing is rendered until the multi printer itself is started.
	multi.Start()

	// Advance the printers at different rates to show that they update
	// independently.
	for i := 1; i <= 100; i++ {
		pb1.Increment()

		if i%2 == 0 {
			pb2.Add(3)
		}

		if i%5 == 0 {
			pb3.Increment()
		}

		if i%10 == 0 {
			pb4.Increment()
		}

		if i%3 == 0 {
			pb5.Increment()
		}

		if i%50 == 0 {
			spinner1.Success("Spinner 1 is done!")
		}

		if i%60 == 0 {
			spinner2.Fail("Spinner 2 failed!")
		}

		time.Sleep(time.Millisecond * 50)
	}

	multi.Stop()
}
