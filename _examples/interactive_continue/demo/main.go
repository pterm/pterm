package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Create an interactive continue prompt with default settings
	// This will pause the program execution until the user presses enter
	// The message displayed is "Press 'Enter' to continue..."
	prompt := pterm.DefaultInteractiveContinue

	// Show the prompt and wait for user input
	// The returned result is the user's input (should be empty as it's a continue prompt)
	// The second return value is an error which is ignored here
	result, _ := prompt.Show()

	// Print a blank line for better readability
	pterm.Println()

	// Print the user's input with an info prefix
	// As this is a continue prompt, the input should be empty
	pterm.Info.Printfln("You answered: %s", result)
}
