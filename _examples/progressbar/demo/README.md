# progressbar/demo

![Animation](animation.svg)

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
