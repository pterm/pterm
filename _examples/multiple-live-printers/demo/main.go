package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	multi := pterm.DefaultMultiPrinter

	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(25).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(50).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	spinner1, _ := pterm.DefaultSpinner.Start("Spinner 1")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(12).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(96).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	multi.Start()

	// Randomly increment progress bars for demo purposes.
	for i := 1; i <= 100; i++ {
		pb1.Increment()

		if i%4 == 0 {
			pb2.Increment()
		}

		if i%2 == 0 {
			pb3.Increment()
		}

		if i%8 == 0 {
			pb4.Increment()
		}

		if i%3 == 0 {
			pb5.Increment()
		}

		if i%50 == 0 {
			spinner1.Stop()
		}

		time.Sleep(time.Millisecond * 100)
	}
	spinner1.Stop()
}
