package main

import "github.com/pterm/pterm"

func main() {
	// Create an interactive text input with a mask for password input
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")

	// Show the password input prompt and store the result
	result, _ := passwordInput.Show("Enter your password")

	// Get the default logger from PTerm
	logger := pterm.DefaultLogger

	// Log the received password (masked)
	// Note: In a real-world application, you should never log passwords
	logger.Info("Password received", logger.Args("password", result))
}
