package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	var counter int
	area, _ := pterm.DefaultArea.WithCallback(func() {
		counter++
	}).Start()
	// counter = 1

	// counter = 1 + 4
	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	// counter = 5 + 1
	area.Update(pterm.Sprintf("Counter callback: %d\n The callback is called before Update the content", counter))
}
