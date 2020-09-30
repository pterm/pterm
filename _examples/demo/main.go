package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Change this to time.Millisecond*200 to speed up the demo.
	// Useful when debugging.
	const second = time.Second

	pterm.Header.SetBackgroundStyle(pterm.BgLightBlue).SetMargin(10).Println("PTDP - PTerm Demo Program")
	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to setup!" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory.")

	time.Sleep(second)
}

func clear() {
	print("\033[H\033[2J")
}
