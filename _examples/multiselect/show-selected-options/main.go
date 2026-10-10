package main

import "github.com/pterm/pterm"

func main() {
	// By default the answered prompt only counts the selection. With
	// WithShowSelectedOptions it also lists what was picked, which reads
	// well in logs and recordings.
	pterm.Multiselect.
		WithOptions([]string{"linux/amd64", "linux/arm64", "darwin/amd64", "darwin/arm64", "windows/amd64"}).
		WithDefaultOptions([]string{"linux/amd64", "darwin/arm64"}).
		WithShowSelectedOptions().
		Show("Build for which platforms?")
}
