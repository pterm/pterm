# progressbar/multiple

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
	"time"
)

func main() {
	multi := pterm.DefaultMultiPrinter

	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	multi.Start()

	// Randomly increment progress bars for demo purposes.
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

		time.Sleep(time.Millisecond * 50)
	}

	multi.Stop()
}

```
