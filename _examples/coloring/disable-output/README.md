# coloring/disable-output

![Animation](https://vhs.charm.sh/vhs-6JuPDJLeFeUCl3biL6qiP0.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// DisableOutput silences all PTerm printers globally until EnableOutput
	// is called. Iterations 5-9 below produce no output at all.
	for i := 0; i < 15; i++ {
		switch i {
		case 5:
			pterm.Info.Println("Disabled Output!")
			pterm.DisableOutput()
		case 10:
			pterm.EnableOutput()
			pterm.Info.Println("Enabled Output!")
		}

		pterm.Printf("Printing something... [%d/%d]\n", i, 15)
	}
}
```
