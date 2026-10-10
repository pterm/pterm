package main

import "github.com/pterm/pterm"

func main() {
	logger := pterm.DefaultLogger

	// Arguments are printed in the order you pass them.
	logger.Info("Request handled", logger.Args("status", 200, "method", "GET", "path", "/users"))

	// WithSortArguments sorts them alphabetically by key instead, which makes
	// lines with the same keys easier to compare.
	sorted := logger.WithSortArguments()
	sorted.Info("Request handled", sorted.Args("status", 200, "method", "GET", "path", "/users"))
}
