package main

import (
	"context"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The spinner stops as canceled when the context ends, which makes
	// timeouts a one-liner.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	spinner, _ := pterm.Spinner.WithContext(ctx).Start("Waiting for the database")

	<-ctx.Done()
	time.Sleep(100 * time.Millisecond)

	// Stopping a spinner that already ended does nothing.
	_ = spinner.Stop()
}
