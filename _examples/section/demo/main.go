package main

import "github.com/pterm/pterm"

func main() {
	// Sections split output into parts. Each one starts with a blank line,
	// so the program does not have to add spacing itself.
	pterm.Section.Println("Installation")
	pterm.Text.Println("Run the installer and follow the prompts.")

	pterm.Section.Println("Configuration")
	pterm.Text.Println("Settings live in ~/.config/mjw/config.toml.")
}
