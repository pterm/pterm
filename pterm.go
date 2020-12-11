// Package pterm is a modern go module to beautify console output.
// It can be used without configuration, but if desired, everything can be customized down to the smallest detail.
// View the animated examples here: https://github.com/pterm/pterm#-examples
package pterm

var (
	// Output completely disables output from pterm if set to false. Can be used in CLI application quiet mode.
	Output = true
	// PrintDebugMessages sets if messages printed by the DebugPrinter should be printed.
	PrintDebugMessages = false
)

// EnableOutput enables the output of PTerm.
func EnableOutput() {
	Output = true
}

// DisableOutput disables the output of PTerm.
func DisableOutput() {
	Output = false
}

// EnableDebugMessages enables the output of debug printers.
func EnableDebugMessages() {
	PrintDebugMessages = true
}

// DisableDebugMessages disables the output of debug printers.
func DisableDebugMessages() {
	PrintDebugMessages = false
}
