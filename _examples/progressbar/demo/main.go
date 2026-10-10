package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	packages := []string{
		"react", "react-dom", "typescript", "vite", "eslint", "prettier",
		"vitest", "zod", "date-fns", "clsx", "tailwindcss", "postcss",
	}

	bar, _ := pterm.Progressbar.WithTotal(len(packages)).Start("Installing packages")

	for i, pkg := range packages {
		bar.UpdateTitle("Installing " + pkg)

		// One package takes a while, which the ETA picks up.
		if i == 6 {
			time.Sleep(1500 * time.Millisecond)
		}

		time.Sleep(300 * time.Millisecond)

		// Output printed while the bar runs appears above it.
		pterm.Success.Println("Installed " + pkg)
		bar.Increment()
	}

	// The bar stopped on its own when it reached its total and left a
	// status line behind.
}
