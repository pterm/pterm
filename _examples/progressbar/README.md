### progressbar/demo

![Animation](https://vhs.charm.sh/vhs-7wWqFXiNFydcIw4FCmNZQG.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Pretend we have a list of packages to download.
var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	p, _ := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		// Simulate one download taking much longer than the rest.
		if i == 6 {
			time.Sleep(time.Second * 3)
		}

		p.UpdateTitle("Downloading " + fakeInstallList[i])

		// Printing through pterm while the progressbar runs places the output
		// above the bar instead of breaking it.
		pterm.Success.Println("Downloading " + fakeInstallList[i])
		p.Increment()
		time.Sleep(time.Millisecond * 350)
	}
}
```

</details>

### progressbar/custom-style

![Animation](https://vhs.charm.sh/vhs-6q2cOGHtVlL0B7qOi8jpZX.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Every visual part of the progressbar can be swapped out. Here the bar
	// gets a retro ASCII look; clearing BarPartialCharacters disables the
	// smooth block-glyph edge, which would clash with plain "#".
	p, _ := pterm.DefaultProgressbar.
		WithTotal(50).
		WithTitle("Installing").
		WithBarCharacter("#").
		WithLastCharacter("#").
		WithBarFiller("-").
		WithBarPartialCharacters(nil).
		WithTitleStyle(pterm.NewStyle(pterm.FgLightYellow)).
		WithBarStyle(pterm.NewStyle(pterm.FgLightMagenta)).
		WithShowElapsedTime(false).
		Start()

	for i := 0; i < p.Total; i++ {
		p.Increment()
		time.Sleep(time.Millisecond * 60)
	}
}
```

</details>

### progressbar/multiple

![Animation](https://vhs.charm.sh/vhs-6DBA2K63ez2hC8yJeWxsAi.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

