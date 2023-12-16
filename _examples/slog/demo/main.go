package main

import (
	"log/slog"

	"github.com/pterm/pterm"
)

func main() {
	// Create a new slog handler with the default PTerm logger
	handler := pterm.NewSlogHandler(&pterm.DefaultLogger)

	// Create a new slog logger with the handler
	logger := slog.New(handler)

	// Log a debug message (won't show by default)
	logger.Debug("This is a debug message that won't show")

	// Change the log level to debug to enable debug messages
	pterm.DefaultLogger.Level = pterm.LogLevelDebug

	// Log a debug message (will show because debug level is enabled)
	logger.Debug("This is a debug message", "changedLevel", true)

	// Log an info message
	logger.Info("This is an info message")

	// Log a warning message
	logger.Warn("This is a warning message")

	// Log an error message
	logger.Error("This is an error message")
}
