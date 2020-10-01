package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var (
	fakeInstallList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
		"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")
)

func main() {
	// Change this to time.Millisecond*200 to speed up the demo.
	// Useful when debugging.
	const second = time.Second

	pterm.Header.SetBackgroundStyle(pterm.BgLightBlue).SetMargin(10).Println("PTDP - PTerm Demo Program")
	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()

	introSpinner := pterm.DefaultSpinner.SetRemoveWhenDone(true).Start("Waiting for 15 seconds...")
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

	clear()

	pterm.Header.SetBackgroundStyle(pterm.BgLightBlue).SetMargin(10).Println("Pseudo Application created with PTerm")

	time.Sleep(second)

	setupSpinner := pterm.DefaultSpinner.Start("Fetching pseudo install list...")
	time.Sleep(second * 4)
	setupSpinner.Success()

	installSpinner := pterm.DefaultSpinner.Start("Installing...")
	for _, s := range fakeInstallList {
		installSpinner.UpdateText("Installing " + s + "...")
		time.Sleep(second)
	}
	installSpinner.Success("Installed all pseudo programs!")

}

func clear() {
	print("\033[H\033[2J")
}
