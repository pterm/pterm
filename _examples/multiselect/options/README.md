# multiselect/options

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Plain strings work as options too. Default options start checked,
	// and are what Show returns when the program does not run in a
	// terminal.
	toppings, _ := pterm.Multiselect.
		WithOptions([]string{"Cheese", "Mushrooms", "Olives", "Onions", "Peppers", "Pineapple"}).
		WithDefaultOptions([]string{"Cheese"}).
		Show("Pick your toppings")

	pterm.Info.Printfln("Ordering a pizza with %d toppings", len(toppings))
}
```
