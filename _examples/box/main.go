package main

import "github.com/pterm/pterm"

func main() {
	panel1 := pterm.DefaultBox.Sprint("Lorem ipsum dolor sit amet, \nconsectetur adipiscing elit, \nsed do eiusmod tempor incididunt \nut labore et dolore \nmagna aliqua.")
	panel2 := pterm.DefaultBox.Sprint("Ut enim ad minim veniam, \nquis nostrud exercitation \nullamco laboris \nnisi ut aliquip \nex ea commodo \nconsequat.")
	panel3 := pterm.DefaultBox.Sprint("Duis aute irure \ndolor in reprehenderit \nin voluptate velit esse cillum \ndolore eu fugiat \nnulla pariatur.")

	panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{panel1}, {panel2}},
		{{panel3}},
	}).Srender()

	pterm.DefaultBox.WithBottomPadding(0).Println(panels)
}
