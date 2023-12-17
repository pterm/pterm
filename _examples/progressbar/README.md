### progressbar/demo

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Slice of strings representing names of pseudo applications to be downloaded.
var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	// Create a progressbar with the total steps equal to the number of items in fakeInstallList.
	// Set the initial title of the progressbar to "Downloading stuff".
	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	// Loop over each item in the fakeInstallList.
	for i := 0; i < p.Total; i++ {
		// Simulate a slow download for the 7th item.
		if i == 6 {
			time.Sleep(time.Second * 3)
		}

		// Update the title of the progressbar with the current item being downloaded.
		p.UpdateTitle("Downloading " + fakeInstallList[i])

		// Print a success message for the current download. This will be printed above the progressbar.
		pterm.Success.Println("Downloading " + fakeInstallList[i])

		// Increment the progressbar by one to indicate progress.
		p.Increment()

		// Pause for 350 milliseconds to simulate the time taken for each download.
		time.Sleep(time.Millisecond * 350)
	}
}

```

</details>

### progressbar/multiple

![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/multiple/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
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

```

</details>

