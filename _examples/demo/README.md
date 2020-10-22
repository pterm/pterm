# demo

![Animation](animation.svg)

```go
package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	introScreen()
	clear()
	pseudoApplicationHeader()
	time.Sleep(second)
	fetchingPseudoInstallList()
	downloadingPseudoList()
	installingPseudoList()
}

func installingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Installing " + pseudoProgramList[i]
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second)
	}
	p.Stop()
}

func downloadingPseudoList() {
	p := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Downloading stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Downloading " + pseudoProgramList[i]
		pterm.Success.Println("Downloading " + pseudoProgramList[i])
		p.Increment()
		time.Sleep(time.Millisecond * 500)
	}
}

func fetchingPseudoInstallList() {
	setupSpinner := pterm.DefaultSpinner.Start("Fetching pseudo install list...")
	time.Sleep(second * 4)
	setupSpinner.Success()
}

func pseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func introScreen() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromTextWithStyle("P", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromTextWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"PTDP - PTerm Demo Program")

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
	time.Sleep(second)
	for i := 14; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(second)
	}
	introSpinner.Stop()
}

func clear() {
	print("\033[H\033[2J")
}

```
