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
