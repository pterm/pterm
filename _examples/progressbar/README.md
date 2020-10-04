# progressbar

![Animation](animation.svg)

```go
package main

import (
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

var vki int

func main() {
	p := pterm.DefaultProgressbar.SetTotal(2000).SetTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		p.Increment()
		time.Sleep(time.Millisecond * 5)
		if i%200 == 0 {
			p.Title = "Downloading " + fakeInstallList[vki]
			vki++
		}
	}
}

```
