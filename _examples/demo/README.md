# demo

![Animation](animation.svg)

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {

	pterm.PrintHeader("You can do many things with PTerm")

	time.Sleep(time.Second * 3)

	pterm.Println(pterm.Cyan("Like this header above!"))

	time.Sleep(time.Second * 2)

	clear()

	pterm.PrintSuccess("You can print success messages!")

	time.Sleep(time.Second * 2)

	pterm.PrintInfo("Or infos!")

	time.Sleep(time.Second)

	pterm.PrintError("Or errors!")

	time.Sleep(time.Second)

	pterm.PrintWarning("Or warnings!")

	time.Sleep(time.Second)

	pterm.PrintDescription("Even descriptions can be printed...")

	time.Sleep(time.Second * 2)

	customPrefixPrinter := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "CUSTOM",
			Style: []pterm.Color{pterm.FgLightRed, pterm.BgBlue},
		},
	}

	customPrefixPrinter.Println("Or a custom PrefixPrinter can be crafted :)")

	time.Sleep(time.Second * 2)

	pterm.Warning.WithScope("custom-scope").Println("PrefixPrinters also support scopes!")

	time.Sleep(time.Second * 4)

	clear()

	pterm.PrintHeader("Everything can be customized!")

	time.Sleep(time.Second * 2)

	headerStyles := []pterm.Style{
		{pterm.BgGreen},
		{pterm.BgWhite},
		{pterm.BgRed},
		{pterm.BgBlue},
		{pterm.BgYellow},
		{pterm.BgLightMagenta},
	}

	for _, style := range headerStyles {
		clear()
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          5,
		}.Println("Everything can be customized")
		time.Sleep(time.Second / 2)
	}

	for i := 0; i < 10; i++ {
		clear()
		style := headerStyles[len(headerStyles)-1]
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          5 + i,
		}.Println("Everything can be customized")
		time.Sleep(time.Millisecond * 100)
	}

	for i := 0; i < 15; i++ {
		clear()
		style := headerStyles[len(headerStyles)-1]
		pterm.HeaderPrinter{
			BackgroundStyle: style,
			TextStyle:       pterm.Style{pterm.FgLightWhite},
			Margin:          15 - i,
		}.Println("Everything can be customized")
		time.Sleep(time.Millisecond * 100)
	}

}

func clear() {
	print("\033[H\033[2J")
}

```
