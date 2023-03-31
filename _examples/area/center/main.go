package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithCenter().Start()

	for i := 0; i < 5; i++ {
		area.Update(pterm.Sprintf("Current count: %d\nAreas can update their content dynamically!", i))
		time.Sleep(time.Second)
	}

	area.Stop()
}
