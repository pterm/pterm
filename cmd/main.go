package main

import (
	"os"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	showcases := []string{
		"Run short time job 1",
		"Run short time job 2",
		"Run short time job 3",
		"Run long time job 4",
	}

	pb, _ := pterm.DefaultProgressbar.
		WithTotal(len(showcases)).
		WithWriter(os.Stdout).
		Start()
	for i, showcase := range showcases {
		pb.UpdateTitle(showcase)
		time.Sleep(1 * time.Second)
		if i >= 3 {
			time.Sleep(3 * time.Second)
		}
		pterm.Success.WithWriter(os.Stdout).Println(showcase)
		pb.Increment()
	}

	time.Sleep(10 * time.Second)
	pb.Stop()
}
