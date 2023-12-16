package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Create a multi printer instance from the default one
	multi := pterm.DefaultMultiPrinter

	// Create five progress bars with a total of 100 units each, and assign each a new writer from the multi printer
	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	// Start the multi printer
	multi.Start()

	// Loop to increment progress bars based on certain conditions
	for i := 1; i <= 100; i++ {
		pb1.Increment() // Increment the first progress bar at each iteration

		if i%2 == 0 {
			pb2.Add(3) // Add 3 units to the second progress bar at every even iteration
		}

		if i%5 == 0 {
			pb3.Increment() // Increment the third progress bar at every fifth iteration
		}

		if i%10 == 0 {
			pb4.Increment() // Increment the fourth progress bar at every tenth iteration
		}

		if i%3 == 0 {
			pb5.Increment() // Increment the fifth progress bar at every third iteration
		}

		time.Sleep(time.Millisecond * 50) // Pause for 50 milliseconds at each iteration
	}

	// Stop the multi printer
	multi.Stop()
}
