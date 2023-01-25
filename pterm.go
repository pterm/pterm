// Package pterm is a modern go module to beautify console output.
// It can be used without configuration, but if desired, everything can be customized down to the smallest detail.
//
// Official docs are available at: https://docs.pterm.sh
//
// View the animated examples here: https://github.com/pterm/pterm#-examples
package pterm

import (
	"github.com/gookit/color"
	"go.uber.org/atomic"
)

var (
	// Output completely disables output from pterm if set to false. Can be used in CLI application quiet mode.
	Output = atomic.NewBool(true)

	// PrintDebugMessages sets if messages printed by the DebugPrinter should be printed.
	PrintDebugMessages = atomic.NewBool(false)

	// RawOutput is set to true if pterm.DisableStyling() was called.
	// The variable indicates that PTerm will not add additional styling to text.
	// Use pterm.DisableStyling() or pterm.EnableStyling() to change this variable.
	RawOutput = atomic.NewBool(false)
)

func init() {
	color.ForceColor()
}

// EnableOutput enables the output of PTerm.
func EnableOutput() {
	Output.Store(true)
}

// DisableOutput disables the output of PTerm.
func DisableOutput() {
	Output.Store(false)
}

// EnableDebugMessages enables the output of debug printers.
func EnableDebugMessages() {
	PrintDebugMessages.Store(true)
}

// DisableDebugMessages disables the output of debug printers.
func DisableDebugMessages() {
	PrintDebugMessages.Store(false)
}

// EnableStyling enables the default PTerm styling.
// This also calls EnableColor.
func EnableStyling() {
	RawOutput.Store(false)
	EnableColor()
}

// DisableStyling sets PTerm to RawOutput mode and disables all of PTerms styling.
// You can use this to print to text files etc.
// This also calls DisableColor.
func DisableStyling() {
	RawOutput.Store(true)
	DisableColor()
}

// RecalculateTerminalSize updates already initialized terminal dimensions. Has to be called after a termina resize to guarantee proper rendering. Applies only to new instances.
func RecalculateTerminalSize() {
	// keep in sync with DefaultBarChart
	DefaultBarChart.Width = GetTerminalWidth() * 2 / 3
	DefaultBarChart.Height = GetTerminalHeight() * 2 / 3
	DefaultParagraph.MaxWidth = GetTerminalWidth()
}
