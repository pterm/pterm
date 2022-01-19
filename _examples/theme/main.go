package main

import (
	"github.com/pterm/pterm"
	"reflect"
)

func main() {
	// Print info.
	pterm.Info.Println("These are the default theme styles.\n" +
		"You can modify them easily to your personal preference,\n" +
		"or create new themes from scratch :)")

	pterm.Println() // Print one line space.

	// Print every value of the default theme with its own style.
	v := reflect.ValueOf(pterm.ThemeDefault)
	typeOfS := v.Type()

	var panelData pterm.Panels

	if typeOfS == reflect.TypeOf(pterm.Theme{}) {
		for i := 0; i < v.NumField(); i++ {
			field, ok := v.Field(i).Interface().(pterm.Style)
			if ok {
				field.Print(typeOfS.Field(i).Name)
				pterm.Print("  ")
				if i%5 == 0 {
					pterm.Println()
				}
			}
		}
	}

	pterm.DefaultPanel.WithPanels(panelData).Render()
}
