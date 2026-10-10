package main

import "github.com/pterm/pterm"

func main() {
	// "y" and "n" answer at once. The arrow keys move between the answers
	// and enter takes the highlighted one, which starts on the default.
	ok, _ := pterm.Confirm.
		WithDefaultValue(true).
		Show("Delete 23 caches?")

	if !ok {
		pterm.Info.Println("Nothing was deleted")
		return
	}

	pterm.Success.Println("Deleted 23 caches")
}
