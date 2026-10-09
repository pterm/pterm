# slog/demo

![Animation](https://vhs.charm.sh/vhs-5McCxcQOHt11nNBW1ZHpgP.gif)

```go
package main

import (
	"log/slog"

	"github.com/pterm/pterm"
)

func main() {
	// PTerm's logger can act as a handler for the standard library's slog
	// package, so existing slog code gets styled output for free.
	handler := pterm.NewSlogHandler(&pterm.DefaultLogger)
	logger := slog.New(handler)

	// The PTerm logger decides which levels are shown. Its default level is
	// Info, so this debug message is dropped.
	logger.Debug("This is a debug message that won't show")

	// Lowering the level on the PTerm logger takes effect immediately.
	pterm.DefaultLogger.Level = pterm.LogLevelDebug

	logger.Debug("This is a debug message", "changedLevel", true)
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
```
