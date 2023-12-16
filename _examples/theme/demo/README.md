# theme/demo

![Animation](animation.svg)

```go
package main

import (
	"github.com/pterm/pterm"
	"reflect"
	"time"
)

func main() {
	// Print an informational message about the default theme styles.
	pterm.Info.Println("These are the default theme styles.\nYou can modify them easily to your personal preference,\nor create new themes from scratch :)")

	// Print a blank line for better readability.
	pterm.Println()

	// Get the value and type of the default theme.
	v := reflect.ValueOf(pterm.ThemeDefault)
	typeOfS := v.Type()

	// Check if the type of the default theme is 'pterm.Theme'.
	if typeOfS == reflect.TypeOf(pterm.Theme{}) {
		// Iterate over each field in the default theme.
		for i := 0; i < v.NumField(); i++ {
			// Try to convert the field to 'pterm.Style'.
			field, ok := v.Field(i).Interface().(pterm.Style)
			if ok {
				// Print the field name using its own style.
				field.Println(typeOfS.Field(i).Name)
			}
			// Pause for a quarter of a second to make the output easier to read.
			time.Sleep(time.Millisecond * 250)
		}
	}
}

```
