package main

import "github.com/pterm/pterm"

func main() {
	panel1, _ := pterm.DefaultBox.WithText("Lorem ipsum dolor sit amet, \nconsectetur adipiscing elit, \nsed do eiusmod tempor incididunt \nut labore et dolore \nmagna aliqua.").Srender()
	panel2, _ := pterm.DefaultBox.WithText("Ut enim ad minim veniam, \nquis nostrud exercitation \nullamco laboris \nnisi ut aliquip \nex ea commodo \nconsequat.").Srender()
	panel3, _ := pterm.DefaultBox.WithText("Duis aute irure \ndolor in reprehenderit \nin voluptate velit esse cillum \ndolore eu fugiat \nnulla pariatur.").Srender()

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{panel1}, {panel2}},
		{{panel3}},
	}).Srender()

	_ = pterm.DefaultBox.WithText(panels).WithBottomPadding(0).Render()
}
