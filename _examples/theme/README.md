### theme/demo

![Animation](https://vhs.charm.sh/vhs-70r74ucxpvyQTPC32H8Cp0.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"reflect"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("These are the default theme styles.\nYou can modify them easily to your personal preference,\nor create new themes from scratch :)")

	pterm.Println()

	// The theme fields are plain pterm.Style values, so we can walk them with
	// reflection and print every style the theme defines without listing each
	// field by hand. Overriding one is as simple as assigning a new Style.
	v := reflect.ValueOf(pterm.ThemeDefault)
	typeOfS := v.Type()

	if typeOfS == reflect.TypeOf(pterm.Theme{}) {
		for i := 0; i < v.NumField(); i++ {
			field, ok := v.Field(i).Interface().(pterm.Style)
			if ok {
				// Print each field name in its own style, so you can see
				// exactly how it looks.
				field.Println(typeOfS.Field(i).Name)
			}

			time.Sleep(time.Millisecond * 250)
		}
	}
}
```

</details>

