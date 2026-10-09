# progressbar/multiple

![Animation](https://vhs.charm.sh/vhs-6DBA2K63ez2hC8yJeWxsAi.gif)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A MultiPrinter renders several live printers at once. Each progressbar
	// writes to its own writer obtained from the multi printer.
	multi := pterm.DefaultMultiPrinter

	pb1, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 1")
	pb2, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 2")
	pb3, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 3")
	pb4, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 4")
	pb5, _ := pterm.DefaultProgressbar.WithTotal(100).WithWriter(multi.NewWriter()).Start("Progressbar 5")

	multi.Start()

	// Advance the bars at different speeds so they visibly run independently.
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
