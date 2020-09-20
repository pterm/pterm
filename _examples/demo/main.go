package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {

	twoSeconds := time.Tick(time.Second * 2)
	// fiveSeconds := time.Tick(time.Second*5)
	// tenSeconds := time.Tick(time.Second*10)
	// second := time.Tick(time.Second)
	// halfSecond := time.Tick(time.Second/2)
	// quarterSecond := time.Tick(time.Second/4)

	pterm.PrintHeader("You can do many things with PTerm")

	<-twoSeconds

	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: []pterm.Color{pterm.FgLightRed, pterm.BgBlue},
		},
	}

	customPrefixPrinter.Println("This is a custom PrefixPrinter :)")

}
