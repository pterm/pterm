package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	flow, _ := pterm.Flow.Start("deploy")

	step := flow.Step("Building image")
	time.Sleep(time.Second)
	step.Success()

	// Any printer started while the flow runs joins its rail. A progress bar
	// animates in place and collapses into a status line when it is done.
	bar, _ := pterm.Progressbar.WithTotal(24).Start("Uploading layers")
	for range 24 {
		time.Sleep(60 * time.Millisecond)
		bar.Increment()
	}

	// Rendered output, such as a table, is indented next to the rail.
	_ = pterm.Table.WithData(pterm.TableData{
		{"Pod", "Status", "Restarts"},
		{"api-7f9c", "running", "0"},
		{"worker-2b1d", "running", "1"},
	}).Render()

	flow.Success("Deployed v2.4.1")
}
