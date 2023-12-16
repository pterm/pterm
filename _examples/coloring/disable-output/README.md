# coloring/disable-output

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Loop from 0 to 14
	for i := 0; i < 15; i++ {
		switch i {
		case 5:
			// At the 5th iteration, print a message and disable the output
			pterm.Info.Println("Disabled Output!")
			pterm.DisableOutput()
		case 10:
			// At the 10th iteration, enable the output and print a message
			pterm.EnableOutput()
			pterm.Info.Println("Enabled Output!")
		}

		// Print a progress message for each iteration
		pterm.Printf("Printing something... [%d/%d]\n", i, 15)
	}
}

```
