package main

import "github.com/pterm/pterm"

func main() {
	// WithCaller adds the file and line that logged, which helps to find
	// where a message comes from in a bigger program.
	logger := pterm.Logger.WithCaller()

	logger.Info("Cache warmed", logger.Args("entries", 2048))
	logger.Warn("Cache almost full", logger.Args("used", "93%"))
}
