package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {

	for _, p := range pterm.AllPrinters {
		p.Println("This is the default ", p.Prefix.Text, " printer.")
		p.WithScope("scope").Println("This is the ", p.Prefix.Text, " printer with a scope.")
		p.WithScope("custom", pterm.New(pterm.FgLightMagenta, pterm.Bold, pterm.BgWhite)).Println("This is the ", p.Prefix.Text, " printer with a custom scope style.")
		time.Sleep(time.Second)
		pterm.Println()
	}

	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: []pterm.Color{pterm.FgLightRed, pterm.BgBlue},
		},
	}

	customPrefixPrinter.Println("This is a custom PrefixPrinter :)")

}
