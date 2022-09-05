package putils

import "github.com/pterm/pterm/internal"

// CenterText returns a centered string with each line centered in respect to the longest line.
func CenterText(text string) string {
	return internal.CenterText(text, 0)
}
