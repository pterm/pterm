package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Fullscreen fills the terminal height and Center puts the content in
	// the middle of it, which suits a countdown or a splash screen.
	area, _ := pterm.Area.WithFullscreen().WithCenter().WithRemoveWhenDone().Start()

	for i := 5; i > 0; i-- {
		area.Update(pterm.TonePrimary.Bold().Sprintf("Launching in %d", i))
		time.Sleep(time.Second)
	}

	area.Stop()
	pterm.Success.Println("Launched")
}
