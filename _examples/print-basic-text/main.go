package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	pterm.DefaultBasic.Println("How do you make something as simple as printing text exciting? 🤔")
	time.Sleep(time.Second * 3)
	pterm.DefaultBasic.Println("By going over every example in excruciating detail, of course! 🤷")
	time.Sleep(time.Second * 3)
	pterm.DefaultBasic.Println("Just Kidding! Here are some examples for your benefit")
	time.Sleep(time.Second * 3)
	pterm.DefaultBasic.Printf("This has been printed using formatting in Printf %d%%\n", 100)
	pterm.DefaultBasic.Println("This has been printed using Println")
	time.Sleep(time.Second * 2)
	pterm.DefaultBasic.Println(" 🎉  Now Styling our examples up a bit 🎉")
	time.Sleep(time.Second * 2)
	pterm.DefaultBasic.WithStyle(pterm.NewStyle(pterm.FgLightCyan)).Println("Thought I would dress up for the occasion!")

}
