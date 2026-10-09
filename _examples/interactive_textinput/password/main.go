package main

import "github.com/pterm/pterm"

func main() {
	// WithMask echoes the given string instead of the typed characters,
	// which turns the text input into a password prompt.
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")

	result, _ := passwordInput.Show("Enter your password")

	// Never log passwords in a real application, this is just a demo.
	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
}
