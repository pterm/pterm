package main

import (
	"github.com/pterm/pterm"
	"log/slog"
)

func main() {
	handler := pterm.NewSlogHandler(&pterm.DefaultLogger)
	logger := slog.New(handler)

	logger.Debug("This is a debug message that won't show")
	pterm.DefaultLogger.Level = pterm.LogLevelDebug // Enable debug messages
	logger.Debug("This is a debug message", "changedLevel", true)
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
