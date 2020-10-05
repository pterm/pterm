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
	p := pterm.DefaultProgressbar.WithTotal(len(fakeInstallList)).WithTitle("Downloading stuff").Start()

	for i := 0; i < p.Total; i++ {
		p.Title = "Downloading " + fakeInstallList[vki]
		pterm.Success.Println("Downloading " + fakeInstallList[vki])
		vki++
		p.Increment()
		time.Sleep(time.Millisecond * 500)
	}

	pterm.Success.Println("Finished downloading!")

	time.Sleep(time.Second * 5)
}

```
